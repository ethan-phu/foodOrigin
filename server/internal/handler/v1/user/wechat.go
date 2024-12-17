// Package user handles user-related operations
package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"knowFood/internal/model"
	"knowFood/utils/config"
	"knowFood/utils/jwt"
	"knowFood/utils/response"
	"knowFood/utils/xerrors"
	"knowFood/utils/xerrors/ecode"
	jtime "knowFood/utils/xtime"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// WechatLoginReq 微信小程序登录请求参数
type WechatLoginReq struct {
	Code     string `json:"code" binding:"required"` // 小程序登录code
	Nickname string `json:"nickName"`                // 用户昵称
	Avatar   string `json:"avatarUrl"`               // 用户头像
}

// WechatLoginResp 微信登录响应
type WechatLoginResp struct {
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符
	ErrCode    int    `json:"errcode"`     // 错误码
	ErrMsg     string `json:"errmsg"`      // 错误信息
}

// WechatLogin 微信小程序登录
func (uh *Handler) WechatLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req WechatLoginReq
		if err := c.ShouldBindJSON(&req); err != nil {
			response.JSON(c, xerrors.WithCode(ecode.ValidateErr, "请求参数无效"), nil)
			return
		}

		// 调用微信登录接口获取openid
		wxResp, err := getWechatOpenID(req.Code)
		if err != nil {
			response.JSON(c, xerrors.Wrap(err, ecode.WechatLoginErr, "微信登录失败"), nil)
			return
		}

		if wxResp.ErrCode != 0 {
			response.JSON(c, xerrors.WithCode(ecode.WechatCodeInvalidErr, fmt.Sprintf("微信登录失败: %s", wxResp.ErrMsg)), nil)
			return
		}

		// 查找或创建用户
		user, err := uh.userSrv.GetOrCreateWechatUser(context.TODO(), &model.User{
			OpenID:   wxResp.OpenID,
			UnionID:  wxResp.UnionID,
			Nickname: req.Nickname,
			Avatar:   req.Avatar,
		})
		if err != nil {
			response.JSON(c, xerrors.Wrap(err, ecode.UserRegisterErr, "用户创建失败"), nil)
			return
		}

		// 生成JWT token
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

// getWechatOpenID 获取微信OpenID
func getWechatOpenID(code string) (*WechatLoginResp, error) {
	if code == "" {
		return nil, xerrors.WithCode(ecode.ValidateErr, "code不能为空")
	}

	// 微信登录凭证校验接口地址
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		config.GlobalConfig.Wechat.AppID,
		config.GlobalConfig.Wechat.AppSecret,
		code,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, xerrors.Wrap(err, ecode.MpApiErr, "调用微信接口失败")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, xerrors.Wrap(err, ecode.MpApiErr, "读取微信接口响应失败")
	}

	var wxResp WechatLoginResp
	if err := json.Unmarshal(body, &wxResp); err != nil {
		return nil, xerrors.Wrap(err, ecode.MpApiErr, "解析微信接口响应失败")
	}

	return &wxResp, nil
}
