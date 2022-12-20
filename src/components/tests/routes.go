package tests

import (
	"github.com/fooins/opsmgt-backend/src/libraries/res"
	"github.com/gin-gonic/gin"
)

// 注册路由
func RouteRegister(group *gin.RouterGroup) {
	// 测试
	group.GET("/test", getTest)
}

// 测试
func getTest(ctx *gin.Context) {
	query := ctx.Request.URL.Query()

	rst, err := Test(query)
	if err != nil {
		res.ResError(ctx, err)
		return
	}

	res.ResSuccess(ctx, rst)
}
