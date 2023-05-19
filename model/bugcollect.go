package model

type BugCollect struct {
	Id    int `gorm:"primary_key;column:id" json:"id"`
	Detail    string `gorm:"column:detail" json:"detail"`
	UpdateTime string `gorm:"column:updateTime" json:"updateTime"`
}

func (d *BugCollect) TableName() string {
	return "bugCollect"
}
