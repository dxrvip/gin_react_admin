package main

import (
	"goVueBlog/models"
	"goVueBlog/routers"
	"goVueBlog/utils"

	"github.com/spf13/viper"
)

func main() {
	// 初始化设置
	utils.InitSetting()
	//初始化数据库
	models.DbInit()
	// 初始化路由
	r := routers.InitUrlsRouter()

	// 运行
	port := viper.GetString("servers.HttpPort")
	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
