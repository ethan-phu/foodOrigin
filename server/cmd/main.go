package main

import (
	"knowFood/internal/middleware"
	"knowFood/internal/middleware/trace"
	"knowFood/internal/repo/mysql"
	"knowFood/server"
	"knowFood/utils/config"
	"knowFood/utils/log"
	"knowFood/utils/version"
)

// 主程序入口
func main() {
	// 解析服务器启动参数
	appOpt := &server.AppOptions{}
	server.ResolveAppOptions(appOpt)
	// 如果需要打印版本信息，则执行打印
	if appOpt.PrintVersion {
		version.PrintVersion()
	}
	// 加载配置文件
	c := config.Load(appOpt.ConfigFilePath)
	// 初始化日志系统
	log.InitLogger(&c.LogConfig,
		log.WithOption("appName", c.AppName),
		log.WithOption("requestId", trace.RequestId()))
	// 确保程序退出时日志能同步
	defer log.Sync()

	// 创建数据库连接
	ds := mysql.NewDefaultMysql(c.DBConfig)
	// 初始化数据库
	ds.Migrate()
	// 创建HTTP服务器实例
	srv := server.NewHttpServer(config.GlobalConfig)
	// 注册服务器关闭时的处理函数，确保数据库连接能被正确关闭
	srv.RegisterOnShutdown(func() {
		if ds != nil {
			// 关闭数据库的连接
			ds.Close()
		}
	})
	// 初始化路由
	router := initRouter(ds)
	// 启动HTTP服务器，应用中间件和路由
	srv.Run(middleware.NewMiddleware(), router)
}
