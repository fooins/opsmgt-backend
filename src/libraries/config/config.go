package config

import (
	"github.com/fooins/opsmgt-backend/src/libraries/env"
	"github.com/fooins/opsmgt-backend/src/libraries/errors"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// 加载配置
func Load() {
	// 设置配置文件名称（不包含扩展名）
	viper.SetConfigName(env.GetEnv())

	// 设置配置文件类型（扩展名）
	viper.SetConfigType("json")

	// 添加配置文件路径（可添加多个）
	viper.AddConfigPath("./config/")

	// 搜索并读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(
			errors.NormalizeError(
				err,
				errors.SetErrorIsTrusted(false),
			),
		)
	}

	// 检查配置项
	validate()
}

// 检查配置项
func validate() {
	// 读取配置
	configServer := ConfigServer{}
	configDb := ConfigDb{}
	configCrypto := ConfigCrypto{}
	configRedis := ConfigRedis{}
	mapstructure.Decode(viper.GetStringMap("server"), &configServer)
	mapstructure.Decode(viper.GetStringMap("db"), &configDb)
	mapstructure.Decode(viper.GetStringMap("crypto"), &configCrypto)
	mapstructure.Decode(viper.GetStringMap("redis"), &configRedis)

	// 执行校验
	err := validator.New().Struct(Config{
		Server: configServer,
		Db:     configDb,
		Crypto: configCrypto,
		Redis:  configRedis,
	})

	// 错误处理
	if err != nil {
		panic(
			errors.NormalizeError(
				err,
				errors.SetErrorIsTrusted(false),
			),
		)
	}
}
