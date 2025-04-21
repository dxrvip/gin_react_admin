package utils

import (
	"github.com/spf13/viper"
)

var (
	JwtKey string
)

func InitSetting() {
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 初始化 JWT Key
	JwtKey = viper.GetString("app.Key")
}
