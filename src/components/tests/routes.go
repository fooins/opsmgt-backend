package tests

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 注册路由
func RouteRegister(group *gin.RouterGroup) {
	// 测试
	group.GET("/test", getTest)
}

// 测试
func getTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"now":   time.Now(),
		"query": ctx.Request.URL.Query(),
	})
}
