package database

import (
	"ledger/conf"
	"log"
)

func DbTableSync() {
	err := conf.Conf.MysqlEngin.Sync(new(User))
	if err != nil {
		log.Fatalf("同步表%s失败: %v\n", UserTableName, err)
	}

	err = conf.Conf.MysqlEngin.Sync(new(Ledger))
	if err != nil {
		log.Fatalf("同步表%s失败: %v\n", LedgerTableName, err)
	}

}
