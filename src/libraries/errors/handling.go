package errors

import (
	"fmt"
	"os"

	"github.com/fooins/opsmgt-backend/src/libraries/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取日志记录器
var logger = log.GetLogger(
	"errorhandling",
	log.SetNoUniFile(true),
	log.SetConsoleAll(true),
	log.SetNoCaller(false),
)

// 处理错误
func HandleError(errorToHandle any, setOpts ...SetHandleOptions) AppError {
	// 异常处理
	defer func() {
		if err := recover(); err != nil {
			// 这里没有记录日志，因为它可能已经失败了
			fmt.Printf("错误处理失败，这是失败信息，以及它试图处理的原始错误信息：%+v %+v", err, errorToHandle)
		}
	}()

	// 默认配置项
	opts := &HandleOptions{
		Fields:     []zap.Field{},
		CallerSkip: 1,
	}

	// 覆盖默认配置项
	for i := range setOpts {
		setOpts[i](opts)
	}

	// 格式化错误对象
	appError := NormalizeError(errorToHandle)

	// 不可信的错误触发服务和进程关闭
	defer func() {
		if !appError.IsTrusted {
			// TODO: https://github.com/gin-gonic/examples/tree/master/graceful-shutdown
			os.Exit(1)
		}
	}()

	// 记录日志
	l := logger.WithOptions(zap.AddCallerSkip(opts.CallerSkip))
	l.Error(appError.Error(), opts.Fields...)
	l.Sync()

	return appError
}

// 处理路由错误
func HandleRouteError(ctx *gin.Context, errorToHandle any) AppError {
	return HandleError(
		errorToHandle,
		SetHandleFields(map[string]string{
			"reqid": ctx.GetString("requestId"),
		}),
		SetHandleCallerSkip(2),
	)
}
