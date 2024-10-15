package main

import (
	"context"
	"goVueBlog/cmd"
	"goVueBlog/globar"
	"goVueBlog/routers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"
)

func main() {

	defer cmd.Clear()
	// 初始化命令行参数
	cmd.Start()

	// 初始化路由
	r := routers.InitUrlsRouter()
	srv := &http.Server{
		Addr:    viper.GetString("servers.HttpPort"),
		Handler: r,
	}
	// 运行服务

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			globar.Logger.Fatal("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	globar.Logger.Infoln("关闭服务器 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		globar.Logger.Fatal("服务器关闭:", err)
	}
	globar.Logger.Infoln("服务器退出")
}
