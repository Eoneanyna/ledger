package main

import (
	"ledger/conf"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	if err := conf.InitConfig(); err != nil {
		panic(err)
	}

	if err := conf.Conf.GinEngine.Run(":" + conf.Conf.Server.Port); err != nil {
		panic(err)
	}

}
