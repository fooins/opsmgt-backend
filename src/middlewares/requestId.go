package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取头信息
		requestId := ctx.GetHeader("FI-Request-ID")

		// 若不存在则创建
		if len(requestId) == 0 || requestId == "" {
			requestId = uuid.NewV4().String()
		}

		// 设置上下文
		ctx.Set("requestId", requestId)
	}
}
