package main

import (
	"github.com/fooins/opsmgt-backend/src/components/tests"
	"github.com/gin-gonic/gin"
)

// 设置路由
func Routeing(app *gin.Engine) {
	// 创建路由组
	group := app.Group("v1.0")

	// 嵌套路由
	tests.RouteRegister(group)
}
