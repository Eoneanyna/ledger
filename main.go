package main

import (
	"ledger/conf"
	"ledger/database"
	"ledger/route"
)

func main() {
	if err := conf.InitConfig(); err != nil {
		panic(err)
	}
	//注册路由
	route.RegisterRoute(conf.Conf.GinEngine)

	//检查表结构，若没有则创建
	database.DbTableSync()

	if err := conf.Conf.GinEngine.Run(":" + conf.Conf.Server.Port); err != nil {
		panic(err)
	}

}
