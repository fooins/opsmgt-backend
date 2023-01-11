package config

// 配置结构
type Config struct {
	// 服务配置
	Server ConfigServer `json:"server" validate:"required"`
	// 数据库配置
	Db ConfigDb `json:"db" validate:"required"`
	// 加密相关配置
	Crypto ConfigCrypto `json:"crypto" validate:"required"`
	// Redis配置
	Redis ConfigRedis `json:"redis" validate:"required"`
}

// 服务配置
type ConfigServer struct {
	// 端口号
	Port uint `json:"port" validate:"gte=0,lte=65535"`
}

// 数据库配置
type ConfigDb struct {
	// 主机名或IP
	Host string `json:"host" validate:"required"`
	// 端口号
	Port uint `json:"port" validate:"gte=0,lte=65535"`
	// 用户名
	Username string `json:"username" validate:"required"`
	// 密码
	Password string `json:"password" validate:"required"`
	// 数据库名
	Database string `json:"database" validate:"required"`
}

// 加密相关配置
type ConfigCrypto struct {
	// AES密钥
	AesKey string `json:"aesKey" validate:"required"`
}

// Redis配置
type ConfigRedis struct {
	// 主机名或IP
	Host string `json:"host" validate:"required"`
	// 端口号
	Port uint `json:"port" validate:"gte=0,lte=65535"`
	// 密码
	Password string `json:"password" validate:"required"`
	// 数据库索引
	Db uint `json:"db" validate:"gte=0"`
}
