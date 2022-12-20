package main

import (
	"github.com/fooins/opsmgt-backend/src/middlewares"
	"github.com/gin-gonic/gin"
)

// 创建一个应用
func NewApp() *gin.Engine {
	// 创建一个新的空白 Gin 实例
	app := gin.New()

	// 错误处理和恢复
	app.Use(middlewares.Recovery())

	// 处理请求ID
	app.Use(middlewares.RequestId())

	// 错误处理和恢复
	app.Use(middlewares.Recovery())

	// 记录访问日志
	app.Use(middlewares.AccessLog())

	// 设置路由
	Routeing(app)

	return app
}
