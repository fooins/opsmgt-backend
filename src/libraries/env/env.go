package env

import (
	"os"
	"strings"
)

// 环境变量值
var goEnv = strings.TrimSpace(os.Getenv("GO_ENV"))

// 程序当前是否运行在生产环境下
func IsProd() bool {
	return goEnv == "production"
}

// 程序当前是否运行在开发环境下
func IsDev() bool {
	return goEnv == "development"
}

// 获取环境变量值
func GetEnv() string {
	return goEnv
}
