package conf

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var Conf config

type config struct {
	ProjectName string
	GinEngine   *gin.Engine
	MysqlEngin  *gorm.DB
	Server      ServerConfig
	Database    DatabaseConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func InitConfig() error {
	viper.SetConfigName("config") // 配置文件名称（不带扩展名）
	viper.SetConfigType("yaml")   // 如果需要，显式指定配置类型
	viper.AddConfigPath("./")     // 添加配置文件搜索路径

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		return err
	}

	//初始化数据库

	//初始化api

	return nil
}
