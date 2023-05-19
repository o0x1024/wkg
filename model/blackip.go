package model

type BlackIpList struct {
	Id    int `gorm:"primary_key;column:id" json:"id"`
	Ip    string `gorm:"column:ip" json:"cid"`
	Count int `gorm:"column:count" json:"count"`
}

func (d *BlackIpList) TableName() string {
	return "blackiplist"
}
