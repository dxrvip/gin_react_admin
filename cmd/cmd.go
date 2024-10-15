package cmd

import (
	"goVueBlog/config"
	"goVueBlog/globar"
	"goVueBlog/utils"
)

func Start() {
	// 初始化配置文件
	utils.InitSetting()
	// 初始化日志
	globar.Logger = config.InitLogger()
	//初始化数据库
	globar.DB = config.DbInit()

}

func Clear() {

}
