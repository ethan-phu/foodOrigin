// Created on 2023/3/4.
// @author tony
// email xmgtony@gmail.com
// description

package router

import (
	"knowFood/internal/handler/v1/user"
	"knowFood/internal/middleware"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	userHandler *user.Handler
}

func NewApiRouter(
	userHandler *user.Handler) *ApiRouter {
	return &ApiRouter{
		userHandler: userHandler,
	}
}

// Load 实现了server/http.go:40
func (ar *ApiRouter) Load(g *gin.Engine) {
	// 登录
	g.POST("/login", ar.userHandler.Login())
	// 注册
	g.POST("/register", ar.userHandler.Register())
	// 微信小程序登录
	g.POST("/wechat/login", ar.userHandler.WechatLogin())
	// user group
	ug := g.Group("/v1/user", middleware.AuthToken())
	{
		ug.GET("", ar.userHandler.GetUserInfo())
		ug.POST("/refresh-token", ar.userHandler.RefreshToken())
	}
}
