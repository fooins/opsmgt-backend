package main

import (
	"github.com/fooins/opsmgt-backend/src/libraries/config"
	"github.com/fooins/opsmgt-backend/src/libraries/errors"
	"github.com/fooins/opsmgt-backend/src/libraries/log"
	"go.uber.org/zap"
)

// 获取日志记录器
var logger = log.GetLogger(
	"main",
	log.SetLevelNotProd(zap.InfoLevel),
	log.SetLevelProd(zap.InfoLevel),
	log.SetNoErrorFile(true),
	log.SetConsoleAll(true),
	log.SetNoCaller(true),
)

func main() {
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			errors.HandleError(
				errors.NormalizeError(
					err,
					errors.SetErrorIsTrusted(false),
				),
			)
		}
	}()

	// 加载配置
	config.Load()
	logger.Info("加载配置成功")

	// 启动 Web 服务
	logger.Info("Web服务启动成功")
	StartHttpServer()
}
