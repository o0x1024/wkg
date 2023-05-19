package model


type BugReport struct {
	Id    int `gorm:"primary_key;column:id" json:"id"`
	Detail    string `gorm:"column:detail" json:"detail"`
	UpdateTime string `gorm:"column:updateTime" json:"updateTime"`
	IsNew *bool `gorm:"column:isNew" json:"isNew"`


}

func (d *BugReport) TableName() string {
	return "bugReport"
}
