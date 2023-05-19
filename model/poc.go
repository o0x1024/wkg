package model

type Poc struct {
	Id         int    `json:"id,omitempty" gorm:"primary_key;column:id"`
	PocName    string `json:"pocName,omitempty" gorm:"column:pocName"`
	PocContent string `json:"pocContent,omitempty" gorm:"column:pocContent"`
	UpdateTime string `json:"updateTime,omitempty" gorm:"column:updateTime"`
}

func (d *Poc) TableName() string {
	return "poc"
}
