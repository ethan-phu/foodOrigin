// author: maxf
// date: 2022-04-01 16:38
// version: 中间件初始化

package middleware

import (
	"knowFood/internal/handler/ping"
	"knowFood/internal/middleware/trace"
	"knowFood/utils/response"
	"knowFood/utils/xerrors"
	"knowFood/utils/xerrors/ecode"

	"github.com/gin-gonic/gin"
)

// middleware 实现Router接口
// 便于服务启动时加载, middleware本质跟handler无区别
type middleware struct {
}

func NewMiddleware() *middleware {
	return &middleware{}
}

// Load 注册中间件和公共路由，注册中间件和公共路由
func (m *middleware) Load(g *gin.Engine) {
	// 注册中间件，use就是使用注册中间件
	g.Use(gin.Recovery())
	g.Use(NoCache())
	g.Use(Options())
	g.Use(Secure())
	g.Use(trace.SetRequestId())
	// 记录每次请求的请求信息和响应信息
	g.Use(ReqLogger())
	// 404
	g.NoRoute(func(c *gin.Context) {
		// 往c上下文里面写入内容：xerrors.WithCode(ecode.NotFoundErr, "404 not found!")
		response.JSON(c, xerrors.WithCode(ecode.NotFoundErr, "404 not found!"), nil)
	})
	// ping server
	g.GET("/ping", ping.Ping())
}
