package middlewares

import (
	"time"

	"github.com/fooins/opsmgt-backend/src/libraries/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取日志记录器
var accessLogger = log.GetLogger(
	"accesslog",
	log.SetLevelNotProd(zap.InfoLevel),
	log.SetLevelProd(zap.InfoLevel),
	log.SetNoErrorFile(true),
	log.SetConsoleNot(true),
	log.SetNoCaller(true),
)

// 记录访问日志
func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 开始时间
		start := time.Now()

		// 请求ID
		reqid := ctx.GetString("requestId")

		// 记录请求日志
		accessLogger.Info(
			"REQ",
			zap.String("reqid", reqid),
			zap.String("ip", ctx.ClientIP()),
			zap.String("method", ctx.Request.Method),
			zap.String("proto", ctx.Request.Proto),
			zap.String("uri", ctx.Request.RequestURI),
			zap.Int64("len", ctx.Request.ContentLength),
			zap.String("ua", ctx.Request.UserAgent()),
		)
		accessLogger.Sync()

		// 处理下游中间件
		ctx.Next()

		// 记录响应日志
		defer func() {
			accessLogger.Info(
				"RES",
				zap.String("reqid", reqid),
				zap.Int64("size", int64(ctx.Writer.Size())),
				zap.Int("status", ctx.Writer.Status()),
				zap.Duration("cos", time.Until(start)),
			)
			accessLogger.Sync()
		}()
	}
}
