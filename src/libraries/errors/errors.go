package errors

import (
	"fmt"
	"net/http"
)

// 错误码
const (
	UNAUTHORIZED         = "UNAUTHORIZED"
	ACCESS_DENIED        = "ACCESS_DENIED"
	INVALID_REQUEST      = "INVALID_REQUEST"
	NOT_FOUND            = "NOT_FOUND"
	INTERNAL_SERVERERROR = "INTERNAL_SERVERERROR"
	SERVICE_UNAVAILABLE  = "SERVICE_UNAVAILABLE"
	GENERAL_EXCEPTION    = "GENERAL_EXCEPTION"
)

// 错误码对应的默认消息
var Message = map[string]string{
	UNAUTHORIZED:         "客户端未经授权",
	ACCESS_DENIED:        "客户端没有执行该操作的权限",
	INVALID_REQUEST:      "请求格式有误或不正确",
	NOT_FOUND:            "所请求的资源不存在",
	INTERNAL_SERVERERROR: "处理请求时出现服务端内部错误",
	SERVICE_UNAVAILABLE:  "该服务暂时不可用。可以过段时间之后再重复该请求",
	GENERAL_EXCEPTION:    "发生未指定错误",
}

// 统一错误类
type AppError struct {
	Code       string `json:"code"`       // 错误代码
	Message    string `json:"message"`    // 错误消息
	HttpStatus int    `json:"httpStatus"` // HTTP 状态代码
	IsTrusted  bool   `json:"isTrusted"`  // 是否可信的错误，不可信的错误通常会触发服务和进程关闭
}

// 实现错误方法
func (appError *AppError) Error() string {
	return fmt.Sprintf(
		"%s (%s - %d - %t)",
		appError.Message,
		appError.Code,
		appError.HttpStatus,
		appError.IsTrusted,
	)
}

// 创建一个统一错误
func New(code string, message string, httpStatus int, isTrusted bool) AppError {
	return AppError{
		Code:       code,
		Message:    message,
		HttpStatus: httpStatus,
		IsTrusted:  isTrusted,
	}
}

// 创建一个默认的统一错误
func NewDefault(message string) AppError {
	return AppError{
		Code:       INTERNAL_SERVERERROR,
		Message:    message,
		HttpStatus: http.StatusInternalServerError,
		IsTrusted:  true,
	}
}
