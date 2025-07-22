package database

import (
	"fmt"
	"ledger/conf"
)

var LedgerTableName = "ledger"

type Ledger struct {
	Id     int64 `json:"id" xorm:"pk autoincr"`
	UserId int64 `json:"user_id"`
	//金额
	Amount int `json:"amount"`
	//来源 支付宝、微信、银行卡等
	AmountFrom string `json:"amount_from"`
	//自定义标签ID
	TagId int64 `json:"tag_id"`
	//时间戳
	Timestamp int64 `json:"timestamp"`
	//描述
	Description string `json:"description"`
}

func GetLedgerList(userId int64, startTimestamp int64, endTimestamp int64, page int, pageSize int) (resp []*Ledger, total int64, err error) {
	var ledgers []*Ledger

	// 计算偏移量
	offset := (page - 1) * pageSize

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 查询数据
	err = conf.Conf.MysqlEngin.Where("user_id = ? AND timestamp >= ? AND timestamp <= ?", userId, startTimestamp, endTimestamp).Limit(pageSize, offset).Find(&ledgers)
	if err != nil {
		return nil, 0, err
	}

	// 获取总记录数
	total, err = conf.Conf.MysqlEngin.Where("user_id = ? AND timestamp >= ? AND timestamp <= ?", userId, startTimestamp, endTimestamp).Count(&Ledger{})
	if err != nil {
		return nil, 0, err
	}

	return ledgers, total, nil
}

func InsertLedger(ledger *Ledger) error {
	_, err := conf.Conf.MysqlEngin.InsertOne(ledger)
	if err != nil {
		return err
	}
	return nil
}

func FindLedger(ledger *Ledger) error {
	has, err := conf.Conf.MysqlEngin.Get(ledger)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("ledger not found")
	}
	return nil
}

func UpdateLedger(ledgerId int64, m map[string]interface{}) error {
	_, err := conf.Conf.MysqlEngin.Where("id = ?", ledgerId).Update(m)
	if err != nil {
		return err
	}
	return nil
}
