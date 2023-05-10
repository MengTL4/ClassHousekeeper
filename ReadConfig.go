package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig() (string, string, string) {
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("json")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}
	cookie := config.GetString("Cookie")
	imprint := config.GetString("Imprint")
	identity := config.GetString("identity")
	return cookie, imprint, identity
}
