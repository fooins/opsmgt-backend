package errors

import "go.uber.org/zap"

// 错误处理配置项
type HandleOptions struct {
	// 指定追加记录的字段
	Fields []zap.Field `json:"fields"`
	// 指定跳过的调用者数量（影响日志记录）
	CallerSkip int `json:"callerSkip"`
}

// 类型：设置错误处理配置项
type SetHandleOptions func(*HandleOptions)

// 设置错误处理配置项 Fields
func SetHandleFields(fields map[string]string) SetHandleOptions {
	return func(opts *HandleOptions) {
		for key, val := range fields {
			opts.Fields = append(opts.Fields, zap.String(key, val))
		}
	}
}

// 设置错误处理配置项 CallerSkip
func SetHandleCallerSkip(callerSkip int) SetHandleOptions {
	return func(opts *HandleOptions) {
		opts.CallerSkip = callerSkip
	}
}

// 错误实例配置项
type ErrorOptions struct {
	Code       string `json:"code"`       // 错误代码
	HttpStatus int    `json:"httpStatus"` // HTTP 状态代码
	IsTrusted  bool   `json:"isTrusted"`  // 是否可信的错误，不可信的错误通常会触发服务和进程关闭
}

// 类型：设置错误实例配置项
type SetErrorOptions func(*ErrorOptions)

// 设置错误实例配置项 Code
func SetErrorCode(code string) SetErrorOptions {
	return func(opts *ErrorOptions) {
		opts.Code = code
	}
}

// 设置错误实例配置项 HttpStatus
func SetErrorHttpStatus(httpStatus int) SetErrorOptions {
	return func(opts *ErrorOptions) {
		opts.HttpStatus = httpStatus
	}
}

// 设置错误实例配置项 IsTrusted
func SetErrorIsTrusted(isTrusted bool) SetErrorOptions {
	return func(opts *ErrorOptions) {
		opts.IsTrusted = isTrusted
	}
}
