// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description 用户信息handler

package user

import (
	"context"
	"knowFood/internal/middleware"
	"knowFood/internal/service"
	"knowFood/utils/response"
	"knowFood/utils/xerrors"
	"knowFood/utils/xerrors/ecode"

	"github.com/gin-gonic/gin"
)

// Handler 用户业务handler
type Handler struct {
	userSrv service.UserService
}

func NewUserHandler(_userSrv service.UserService) *Handler {
	return &Handler{
		userSrv: _userSrv,
	}
}

func (uh *Handler) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := middleware.GetUserId(c)
		user, err := uh.userSrv.GetById(context.TODO(), uid)
		if err != nil {
			response.JSON(c, xerrors.Wrap(err, ecode.NotFoundErr, "用户信息为空"), nil)
		} else {
			response.JSON(c, nil, user)
		}
	}
}

// 登出
func (uh *Handler) LogOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		// uid := middleware.GetUserId(c)

	}
}
