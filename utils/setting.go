package utils

import (
	"github.com/spf13/viper"
)

func InitSetting() {
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
