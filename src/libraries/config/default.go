package config

import "github.com/spf13/viper"

// 设置默认配置
func setDefaults() {
	// 服务配置
	viper.SetDefault("server", map[string]any{
		// 端口号
		"port": 8080,
	})

	// 数据库
	viper.SetDefault("db", map[string]any{
		// 主机名或IP
		"host": "127.0.0.1",
		// 端口号
		"port": 3306,
		// 用户名
		"username": "",
		// 密码
		"password": "",
		// 数据库名
		"database": "opsmgt",
	})

	// 加密相关配置
	viper.SetDefault("crypto", map[string]any{
		// AES 密钥
		"aesKey": "zmb6ja36v6q45ejob1z2upbxy5qd5c9i",
	})

	// Redis
	viper.SetDefault("redis", map[string]any{
		// 主机名
		"host": "127.0.0.1",
		// 端口号
		"port": 6379,
		// 密码
		"password": "123456",
		// 数据库索引
		"db": 0,
	})
}
