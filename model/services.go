package model

type Services struct {
	Id         int     `json:"id" gorm:"primary_key;column:id"`
	Cid        int     `json:"cid" gorm:"primary_key;column:cid"`
	Service    string  `json:"service" gorm:"column:service"`
	Host       string  `json:"host" gorm:"column:host"`
	Port       string  `json:"port" gorm:"column:port"`
	Product    string  `json:"product" gorm:"column:product"`
	UpdateTime string  `json:"updateTime" gorm:"column:updateTime"`
	IsNew      *bool   `json:"isNew" gorm:"column:isNew" json:"isNew"`
	Companys   Company `gorm:"ForeignKey:Id;AssociationForeignKey:Cid"`
}

func (d *Services) TableName() string {
	return "services"
}
