package database

import "ledger/conf"

var UserTableName = "user"

type User struct {
	Id       int    `json:"id" xorm:"pk autoincr"` //主键+自增
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Other1   string `json:"other1"` //保留字段
}

func InsertUser(user User) error {
	_, err := conf.Conf.MysqlEngin.InsertOne(user)
	return err
}

func GetUserByName(name string) (User, error) {
	var user User
	_, err := conf.Conf.MysqlEngin.Where("name = ?", name).Get(&user)
	return user, err
}
