package conf

import (
	"github.com/Augenblick-tech/bilibot/lib/db"
	"github.com/spf13/viper"
)

type config struct {
	Server Server
	User   User
	DB     DB
	JWT    JWT
}

type Server struct {
	Addr   string `toml:"addr"`   // 后端运行地址
	Domain string `toml:"domain"` // 指定 swagger 的 base url
}

type User struct {
	LisenInterval int // 监听动态的时间间隔
}

type DB struct {
	DbType db.DbType `mapstructure:"type"` // 1: sqlite3, 2: mysql
	Data   string    `toml:"data"` // 打开数据库所需的字符串信息
}

type JWT struct {
	Secret string `toml:"secret"` // jwt 加密密钥
}

var C config

func LoadDefaultConfig() {
	loadConfig("config", "toml", "./conf")
}

func loadConfig(name, typ, path string) {
	viper.SetConfigName(name)
	viper.SetConfigType(typ)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
}
