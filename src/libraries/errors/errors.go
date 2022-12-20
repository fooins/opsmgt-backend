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
var Messages = map[string]string{
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
func New(message string, setOpts ...SetErrorOptions) AppError {
	// 默认配置项
	opts := &ErrorOptions{
		Code:       INTERNAL_SERVERERROR,
		HttpStatus: http.StatusInternalServerError,
		IsTrusted:  true,
	}

	// 覆盖默认配置项
	for i := range setOpts {
		setOpts[i](opts)
	}

	return AppError{
		Code:       opts.Code,
		Message:    message,
		HttpStatus: opts.HttpStatus,
		IsTrusted:  opts.IsTrusted,
	}
}

// 格式化错误对象
func NormalizeError(errorToHandle any, setOpts ...SetErrorOptions) AppError {
	// 如果是统一错误则直接返回
	if appError, ok := errorToHandle.(AppError); ok {
		return appError
	}

	// 如果实现了 error 接口
	if err, ok := errorToHandle.(error); ok {
		return New(err.Error(), setOpts...)
	}

	// 如果是字符串
	if str, ok := errorToHandle.(string); ok {
		return New(str, setOpts...)
	}

	return New(fmt.Sprintf("错误处理程序收到一个未知的错误类型实例：%+v", errorToHandle), setOpts...)
}
