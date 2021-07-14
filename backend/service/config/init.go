package config

import (
	"github.com/spf13/viper"
)

var configService *service

type service struct {
	Config
}

func Init() {
	configService = new(service)
	viper.AddConfigPath(".")// 设置配置文件和可执行二进制文件在用一个目录
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&configService.Config)
	if err != nil {
		panic(err)
	}

}