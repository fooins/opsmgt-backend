package log

import (
	"go.uber.org/zap/zapcore"
)

// 日志记录器配置项
type Options struct {
	// 指定的日志级别（生产环境）
	LevelProd zapcore.Level `json:"levelProd"`
	// 指定的日志级别（非生产环境）
	LevelNotProd zapcore.Level `json:"levelNotProd"`
	// 不创建统一日志文件（包含所有级别的内容）
	NoUniFile bool `json:"noUniFile"`
	// 不创建错误日志文件（仅包含错误以上级别的内容）
	NoErrorFile bool `json:"noErrorFile"`
	// 不打印到控制台
	ConsoleNot bool `json:"consoleNot"`
	// 打印到控制台
	ConsoleAll bool `json:"consoleAll"`
	// 仅生产环境打印到控制台
	ConsoleProd bool `json:"consoleProd"`
	// 仅非生产环境打印到控制台
	ConsoleNotProd bool `json:"consoleNotProd"`
	// 不记录调用者信息
	NoCaller bool `json:"noCaller"`
}

// 类型：设置选项
type SetOption func(*Options)

// 设置 LevelProd
func SetLevelProd(lvl zapcore.Level) SetOption {
	return func(opts *Options) {
		opts.LevelProd = lvl
	}
}

// 设置 LevelNotProd
func SetLevelNotProd(lvl zapcore.Level) SetOption {
	return func(opts *Options) {
		opts.LevelNotProd = lvl
	}
}

// 设置 NoUniFile
func SetNoUniFile(b bool) SetOption {
	return func(opts *Options) {
		opts.NoUniFile = b
	}
}

// 设置 NoErrorFile
func SetNoErrorFile(b bool) SetOption {
	return func(opts *Options) {
		opts.NoErrorFile = b
	}
}

// 设置 ConsoleNot
func SetConsoleNot(b bool) SetOption {
	return func(opts *Options) {
		opts.ConsoleNot = b
	}
}

// 设置 ConsoleAll
func SetConsoleAll(b bool) SetOption {
	return func(opts *Options) {
		opts.ConsoleAll = b
	}
}

// 设置 ConsoleProd
func SetConsoleProd(b bool) SetOption {
	return func(opts *Options) {
		opts.ConsoleProd = b
	}
}

// 设置 ConsoleNotProd
func SetConsoleNotProd(b bool) SetOption {
	return func(opts *Options) {
		opts.ConsoleNotProd = b
	}
}

// 设置 NoCaller
func SetNoCaller(b bool) SetOption {
	return func(opts *Options) {
		opts.NoCaller = b
	}
}
