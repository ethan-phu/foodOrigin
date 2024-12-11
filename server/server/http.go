// author: xmgtony
// date: 2020-09-28 16:10
// version: 1.0
// 通用server启动类，完成配置加载，缓存初始化，数据库初始化，启动参数绑定

package server

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"knowFood/utils/config"
	"knowFood/utils/log"
	"knowFood/utils/validator"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// HttpServer 代表当前服务端实例
type HttpServer struct {
	config *config.Config
	f      func()
}

// NewHttpServer 创建server实例
func NewHttpServer(config *config.Config) *HttpServer {
	return &HttpServer{
		config: config,
	}
}

// Router 加载路由，使用侧提供接口，实现侧需要实现该接口
type Router interface {
	Load(engine *gin.Engine)
}

// AppOptions 用来接收应用启动时指定的参数
type AppOptions struct {
	PrintVersion   bool   // 打印版本
	ConfigFilePath string // 配置文件路径
}

// ResolveAppOptions 解析启动参数，解析命令行参数
func ResolveAppOptions(opt *AppOptions) {
	var printVersion bool
	var configFilePath string
	flag.BoolVar(&printVersion,
		"v",
		false,
		"-v 选项用于控制是否打印当前项目的版本",
	)
	flag.StringVar(&configFilePath,
		"c", "",
		"-c 选项用于指定要使用的配置文件")
	flag.Parse()

	opt.PrintVersion = printVersion
	opt.ConfigFilePath = configFilePath
}

// Run server的启动入口
// 加载路由, 启动服务
func (s *HttpServer) Run(rs ...Router) {
	var wg sync.WaitGroup
	wg.Add(1)
	// 设置gin启动模式，必须在创建gin实例之前
	gin.SetMode(s.config.Mode)
	g := gin.New()
	// 路由加载
	s.routerLoad(g, rs...)
	// gin validator替换
	validator.LazyInitGinValidator(s.config.Language)

	// health check
	go func() {
		if err := Ping(s.config.Port, s.config.MaxPingCount); err != nil {
			log.Fatal("server no response")
		}
		log.Info("server started success!", "server at http://localhost"+s.config.Port)
	}()

	srv := http.Server{
		Addr:    s.config.Port,
		Handler: g,
	}
	if s.f != nil {
		// 注册http服务关闭的函数
		srv.RegisterOnShutdown(s.f)
	}
	// graceful shutdown
	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	go func() {
		<-sgn
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Error("server shutdown error", "error", err)
		}
		wg.Done()
	}()

	err := srv.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Error("server start failed!", "port", s.config.Port)
			return
		}
	}

	// 等待资源关闭
	wg.Wait()
	log.Info("server stop on port!", "port", s.config.Port)
}

// RouterLoad 加载自定义路由
func (s *HttpServer) routerLoad(g *gin.Engine, rs ...Router) *HttpServer {
	for _, r := range rs {
		r.Load(g)
	}
	return s
}

// RegisterOnShutdown 注册shutdown后的回调处理函数，用于清理资源
func (s *HttpServer) RegisterOnShutdown(_f func()) {
	s.f = _f
}

// Ping 用来检查是否程序正常启动
func Ping(port string, maxCount int) error {
	seconds := 1
	if len(port) == 0 {
		panic("Please specify the service port")
	}
	if !strings.HasPrefix(port, ":") {
		port += ":"
	}
	url := fmt.Sprintf("http://localhost%s/ping", port)
	for i := 0; i < maxCount; i++ {
		resp, err := http.Get(url)
		if nil == err && resp != nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		// 服务不在线，需要等待，重试多次
		log.Info(fmt.Sprintf("等待服务在线, 已等待 %d 秒，最多等待 %d 秒", seconds, maxCount))
		time.Sleep(time.Second * 1)
		seconds++
	}
	return fmt.Errorf("服务启动失败，端口 %s", port)
}
