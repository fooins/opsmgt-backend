package log

import (
	"fmt"
	"os"

	"github.com/fooins/opsmgt-backend/src/libraries/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// 缓存日志记录器
var loggers = make(map[string]*zap.Logger)

// 日志格式配置
var encoderConfig = zapcore.EncoderConfig{
	MessageKey:    "msg",
	LevelKey:      "lvl",
	TimeKey:       "time",
	NameKey:       "name",
	CallerKey:     "caller",
	FunctionKey:   "func",
	StacktraceKey: "stack",

	SkipLineEnding: false,
	LineEnding:     zapcore.DefaultLineEnding,

	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.MillisDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
	EncodeName:     zapcore.FullNameEncoder,

	ConsoleSeparator: "|",
}

// 创建日志记录器
func newLogger(name string, opts *Options) *zap.Logger {
	logger, exists := loggers[name]
	if exists {
		return logger
	}

	var level zap.AtomicLevel
	var cores []zapcore.Core
	var zapopts []zap.Option

	// 日志级别
	if env.IsProd() {
		level = zap.NewAtomicLevelAt(opts.LevelProd)
	} else {
		level = zap.NewAtomicLevelAt(opts.LevelNotProd)
	}

	// 所有级别的日志记录到一个统一的文件中
	// 按照日期滚动
	if !opts.NoUniFile {
		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(&lumberjack.Logger{
				// 文件名
				Filename: fmt.Sprintf("./logs/%s.log", name),
				// 滚动前文件的最大大小（以兆字节为单位）
				// 默认为 100 兆
				MaxSize: 100,
				// 保留旧日志文件的最大天数
				// 默认不根据天数删除旧日志文件
				MaxAge: 30,
				// 保留旧日志文件的最大数量
				// 默认是保留所有旧的日志文件
				MaxBackups: 100,
				// 用于格式化备份文件中的时间戳的时间是否为计算机的本地时间
				// 默认使用 UTC 时间
				LocalTime: true,
				// 旋转的日志文件是否应该使用 gzip 进行压缩
				// 默认情况下不执行压缩
				Compress: true,
			}),
			level,
		))
	}

	// 错误级别的日志单独再记录一份，以便处理
	// 按大小滚动
	if !opts.NoErrorFile {
		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(&lumberjack.Logger{
				// 文件名
				Filename: fmt.Sprintf("./logs/%s-error.log", name),
				// 滚动前文件的最大大小（以兆字节为单位）
				// 默认为 100 兆
				MaxSize: 100,
				// 保留旧日志文件的最大天数
				// 默认不根据天数删除旧日志文件
				MaxAge: 30,
				// 保留旧日志文件的最大数量
				// 默认是保留所有旧的日志文件
				MaxBackups: 100,
				// 用于格式化备份文件中的时间戳的时间是否为计算机的本地时间
				// 默认使用 UTC 时间
				LocalTime: true,
				// 旋转的日志文件是否应该使用 gzip 进行压缩
				// 默认情况下不执行压缩
				Compress: true,
			}),
			zap.NewAtomicLevelAt(zap.ErrorLevel),
		))
	}

	// 打印到控制台
	isConsole := env.IsProd()
	if opts.ConsoleAll {
		isConsole = true
	} else if opts.ConsoleNot {
		isConsole = false
	} else if opts.ConsoleProd {
		isConsole = env.IsProd()
	} else if opts.ConsoleNotProd {
		isConsole = !env.IsProd()
	}
	if isConsole {
		cores = append(cores, zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			level,
		))
	}

	// 添加堆栈跟踪信息
	zapopts = append(zapopts, zap.AddStacktrace(zap.ErrorLevel))
	// 添加进程ID
	zapopts = append(zapopts, zap.Fields(zap.Int("pid", os.Getpid())))
	// 添加调用者信息
	if !opts.NoCaller {
		zapopts = append(zapopts, zap.AddCaller())
	}

	// 创建日志记录器
	logger = zap.New(
		zapcore.NewTee(cores...),
		zapopts...,
	)

	loggers[name] = logger
	return logger
}

// 获取包含缺省配置的日志记录器
func GetDefault(name string) *zap.Logger {
	return newLogger(name, &Options{
		LevelProd:      zap.WarnLevel,
		LevelNotProd:   zap.DebugLevel,
		NoUniFile:      false,
		NoErrorFile:    false,
		ConsoleNot:     true,
		ConsoleAll:     false,
		ConsoleProd:    false,
		ConsoleNotProd: false,
		NoCaller:       false,
	})
}

// 获取日志记录器
func Get(name string, setOpts ...SetOption) *zap.Logger {
	// 默认配置项
	opts := &Options{
		LevelProd:      zap.WarnLevel,
		LevelNotProd:   zap.DebugLevel,
		NoUniFile:      false,
		NoErrorFile:    false,
		ConsoleNot:     true,
		ConsoleAll:     false,
		ConsoleProd:    false,
		ConsoleNotProd: false,
		NoCaller:       false,
	}

	// 覆盖默认配置项
	for i := range setOpts {
		setOpts[i](opts)
	}

	return newLogger(name, opts)
}
