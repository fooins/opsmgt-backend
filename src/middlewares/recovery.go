package middlewares

import (
	"github.com/fooins/opsmgt-backend/src/libraries/errors"
	"github.com/fooins/opsmgt-backend/src/libraries/res"
	"github.com/gin-gonic/gin"
)

// 错误处理和恢复
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 捕捉错误
		defer func() {
			if err := recover(); err != nil {
				// 错误处理
				if appErr, ok := err.(errors.AppError); !ok || appErr.HttpStatus >= 500 {
					errors.HandleRouteError(ctx, err)
				}

				// 响应错误
				res.ResError(ctx, err)
			}
		}()

		// 处理下游中间件
		ctx.Next()
	}
}
