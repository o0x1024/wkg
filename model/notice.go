package model

type Notice struct {
	Id         int    `json:"id,omitempty" gorm:"primary_key;column:id"`
	Cid        int    `json:"cid,omitempty" gorm:"foreignkey;column:cid"`
	Title      string `json:"title,omitempty" gorm:"column:title"`
	PocContent string `json:"content,omitempty" gorm:"column:content"`
	UpdateTime string `json:"updateTime,omitempty" gorm:"column:updateTime"`
}

func (d *Notice) TableName() string {
	return "notice"
}
