package database

import "ledger/conf"

var LedgerTagTableName = "ledger_tag"

// tag可以缓存到redis
type LedgerTag struct {
	TagId    int64  `json:"tag_id" xorm:"pk autoincr" `
	UserId   int64  `json:"user_id"`   //非必须
	TagName  string `json:"tag_name"`  //tag名称
	TagType  int    `json:"tag_type"`  //1.通用tag 2.用户tag（限制创建数量10个）
	TagTopic int    `json:"tag_topic"` //1.生活 2.工作 3.学习 4.运动 5.旅行 6.其他
}

func GetTagList(userId int64) ([]LedgerTag, error) {
	var tagList []LedgerTag
	err := conf.Conf.MysqlEngin.Table(LedgerTagTableName).Where("user_id = ? or user_id = 0", userId).Find(&tagList)
	return tagList, err
}

func GetTagCount(userId int64) (int64, error) {
	return conf.Conf.MysqlEngin.Table(LedgerTagTableName).Where("user_id = ?", userId).Count()
}

func CreateTag(data *LedgerTag) error {
	_, err := conf.Conf.MysqlEngin.Table(LedgerTagTableName).InsertOne(data)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLedgerTag(tagId int64, data LedgerTag) error {
	_, err := conf.Conf.MysqlEngin.Table(LedgerTagTableName).Where("tag_id = ?", tagId).Update(data)
	if err != nil {
		return err
	}
	return nil
}
