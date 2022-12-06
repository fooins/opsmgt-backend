package main

import (
	"github.com/gin-gonic/gin"
)

// 创建一个应用
func NewApp() *gin.Engine {
	// 创建一个新的空白 Gin 实例
	app := gin.New()

	// 设置路由
	Routeing(app)

	return app
}
