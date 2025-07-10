package main

import (
	"ledger/conf"
	"ledger/database"
	"ledger/route"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	if err := conf.InitConfig(); err != nil {
		panic(err)
	}
	//注册路由
	route.RegisterRoute(conf.Conf.GinEngine)

	//检查表结构，若没有则创建
	database.DatabaseTableSync()

	if err := conf.Conf.GinEngine.Run(":" + conf.Conf.Server.Port); err != nil {
		panic(err)
	}

}
