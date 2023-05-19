package model

type Users struct {
	Id         int     `json:"id" gorm:"primary_key;column:id"`
	Username   string     `json:"username" gorm:"column:username"`
	Password   string  `json:"password" gorm:"column:password"`
	CreateTime string  `json:"createTime" gorm:"column:createTime"`
}


func (d *Users)TableName() string {
	return "users"
}