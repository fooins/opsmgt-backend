package tests

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 注册路由
func RouteRegister(group *gin.RouterGroup) {
	// 测试
	group.GET("/test", GetTest)
}

// 测试
func GetTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"now":   time.Now(),
		"query": ctx.Request.URL.Query(),
	})
}
