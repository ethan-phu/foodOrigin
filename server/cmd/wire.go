// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description 使用Google依赖注入工具wire

//go:build wireinject
// +build wireinject

package main

import (
	handlerV1 "knowFood/internal/handler/v1"
	"knowFood/internal/repo/mysql"
	"knowFood/internal/router"
	"knowFood/internal/service"
	"knowFood/server"
	"knowFood/utils/db"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	mysql.ProviderSet,
	service.ProviderSet,
	handlerV1.ProviderSet,
)

// initRouter initializes and returns the router for the server
func initRouter(ds db.IDataSource) server.Router {
	panic(wire.Build(
		providerSet,
		router.ApiRouterProviderSet,
	))
}
