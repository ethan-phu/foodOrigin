// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description

package router

import (
	"knowFood/server"

	"github.com/google/wire"
)

// 提供者
var ApiRouterProviderSet = wire.NewSet(
	NewApiRouter,
	// 注入器，将接口和接口的指针实例对象绑定
	wire.Bind(new(server.Router), new(*ApiRouter)),
)
