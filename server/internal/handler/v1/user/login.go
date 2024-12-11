// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description 用户登录

package user

import (
	"context"
	"knowFood/internal/model"
	"knowFood/tools/security"
	"knowFood/utils/config"
	"knowFood/utils/jwt"
	"knowFood/utils/response"
	"knowFood/utils/xerrors"
	"knowFood/utils/xerrors/ecode"
	jtime "knowFood/utils/xtime"
	"time"

	"github.com/gin-gonic/gin"
)

func (uh *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginReqParam := model.LoginReq{}
		if err := c.ShouldBind(&loginReqParam); err != nil {
			response.JSON(c, xerrors.WithCode(ecode.ValidateErr, err.Error()), nil)
			return
		}
		// 查询用户信息
		user, err := uh.userSrv.GetByMobile(context.TODO(), loginReqParam.Mobile)
		if err != nil {
			response.JSON(c, xerrors.Wrap(err, ecode.UserLoginErr, "登录失败，用户不存在"), nil)
			return
		}

		if !security.ValidatePassword(loginReqParam.Password, user.Password) {
			response.JSON(c, xerrors.WithCode(ecode.UserLoginErr, "登录失败，用户名、密码不匹配"), nil)
			return
		}
		// 生成jwt token
		expireAt := time.Now().Add(24 * 7 * time.Hour)
		claims := jwt.BuildClaims(expireAt, user.Id)
		token, err := jwt.GenToken(claims, config.GlobalConfig.JwtSecret)
		if err != nil {
			response.JSON(c, xerrors.Wrap(err, ecode.UserLoginErr, "生成用户授权token失败"), nil)
			return
		}
		response.JSON(c, nil, struct {
			Token    string     `json:"token"`
			ExpireAt jtime.Time `json:"expire_at"`
		}{
			Token:    token,
			ExpireAt: jtime.Time(expireAt),
		})
	}
}

func (uh *Handler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		registerReq := model.RegisterReq{}
		if err := c.ShouldBindJSON(&registerReq); err != nil {
			response.JSON(c, xerrors.WithCode(ecode.ValidateErr, err.Error()), nil)
			return
		}
		// 查询用户信息, 暂时先用手机号唯一，后面完善之后再定
		_, err := uh.userSrv.GetByMobile(context.TODO(), registerReq.Mobile)
		if err == nil {
			response.JSON(c, xerrors.Wrap(err, ecode.UserRegisterErr, "注册失败，用户已存在"), nil)
			return
		}
		ciphertext, err := security.Encrypt(registerReq.Password)
		if err != nil {
			response.JSON(c, xerrors.Wrap(err, ecode.ValidateErr, "密码计算失败"), nil)
			return
		}
		userInfo := registerReq.ToUserModel(ciphertext)
		err = uh.userSrv.Register(c, &userInfo)
		if err != nil {
			response.JSON(c, xerrors.Wrap(err, ecode.RecordCreateErr, err.Error()), nil)
			return
		}
		response.JSON(c, nil, nil)
	}
}
