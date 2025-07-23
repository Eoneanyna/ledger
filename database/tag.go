package database

var LedgerTagTableName = "ledger_tag"

// tag可以缓存到redis
type LedgerTag struct {
	TagId    int64 `json:"tag_id" xorm:"pk autoincr" `
	UserId   int64 `json:"user_id"`    //非必须
	TagName  int64 `json:"tag_name"`   //tag名称
	TagType  int   `json:"tag_type"`   //1.通用tag 2.用户tag（限制创建数量10个）
	PreTagId int64 `json:"pre_tag_id"` //父级tagId
}
