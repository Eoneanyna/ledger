package conf

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"ledger/auth"
	"log"
	"xorm.io/xorm"
)

var Conf config

type config struct {
	ProjectName string
	GinEngine   *gin.Engine
	MysqlEngin  *xorm.Engine
	Server      ServerConfig
	Database    DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DBName   string `json:"DBName"`
}

func InitConfig() error {
	viper.SetConfigName("config") // 配置文件名称（不带扩展名）
	viper.SetConfigType("yaml")   // 如果需要，显式指定配置类型
	viper.AddConfigPath("./conf") // 添加配置文件搜索路径

	var err error
	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	if err = viper.Unmarshal(&Conf); err != nil {
		return err
	}

	//初始化数据库
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", Conf.Database.UserName, Conf.Database.Password, Conf.Database.IP, Conf.Database.Port, Conf.Database.DBName)

	Conf.MysqlEngin, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("连接数据库失败: %v\n", err)
	}

	//初始化api
	Conf.GinEngine = gin.Default()
	// 使用自定义鉴权中间件
	Conf.GinEngine.Use(auth.AuthMiddleware())

	return nil
}
