package res

import (
	"net/http"

	"github.com/fooins/opsmgt-backend/src/libraries/errors"
	"github.com/gin-gonic/gin"
)

// 响应接口
func Respond(ctx *gin.Context, httpStatus int, code string, message string, data map[string]any) {
	// 设置请求ID响应头
	if reqid := ctx.GetString("requestId"); len(reqid) > 0 {
		ctx.Header("FI-Request-ID", reqid)
	}

	ctx.JSON(httpStatus, map[string]any{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

// 响应成功
func ResSuccess(ctx *gin.Context, message string, data map[string]any) {
	Respond(ctx, http.StatusOK, "SUCCESS", message, data)
}

// 响应错误并携带数据
func ResErrorWithData(ctx *gin.Context, err any, data map[string]any) {
	appErr := errors.NormalizeError(err)
	Respond(ctx, appErr.HttpStatus, appErr.Code, appErr.Message, data)
}

// 响应错误
func ResError(ctx *gin.Context, err any) {
	ResErrorWithData(ctx, err, map[string]any{})
}
