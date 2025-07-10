package database

var UserTableName = "user"

type User struct {
	Id       int    `json:"id" xorm:"pk autoincr"` //主键+自增
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}
