package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// 启动 HTTP 服务
func StartHttpServer() {
	// 创建应用
	app := NewApp()

	// 启动服务
	err := app.Run(fmt.Sprintf(":%d", viper.GetInt("server.port")))
	if err != nil {
		panic(err)
	}
}
