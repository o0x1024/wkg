package model


type MiniProgram struct {
	Id         int     `json:"id" gorm:"primary_key;column:id"`
	Cid        int     `json:"cid" gorm:"foreignkey;column:cid"`
	Name       string  `json:"name" gorm:"column:name"`
	Remarks     string  `json:"remarks" gorm:"column:remarks"`
	UpdateTime string  `json:"updateTime" gorm:"column:updateTime"`
	IsNew		*bool	`json:"isNew" gorm:"column:isNew" json:"isNew"`
}

func (d *MiniProgram)TableName() string {
	return "miniprogram"
}





