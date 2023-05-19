package model

type Task struct {
	TaskId               string `json:"taskId,omitempty" gorm:"primary_key;column:taskId"`
	TaskStatus           *bool  `json:"taskStatus,omitempty" gorm:"primary_key;column:taskStatus"`
	TaskName             string `json:"taskName,omitempty" gorm:"column:taskName"`
	TaskType             string `json:"taskType,omitempty" gorm:"column:taskType"`
	Rate                 int    `json:"rate,omitempty" gorm:"column:rate"`
	CompanyRange         string `json:"companyRange,omitempty" gorm:"column:companyRange"`
	CompanyId            string `json:"companyId,omitempty" gorm:"column:companyId"`
	ComanyName           string `json:"comanyName,omitempty" gorm:"column:comanyName"`
	PocId                int    `json:"pocId,omitempty" gorm:"column:pocId"`
	PocName              string `json:"pocName,omitempty" gorm:"column:pocName"`
	AssetRange           string `json:"assetRange,omitempty" gorm:"column:assetRange"`
	DomainCollectionType string `json:"domainCollectionType,omitempty" gorm:"column:domainCollectionType"`
	PortRange            string `json:"portRange,omitempty" gorm:"column:portRange"`
	UpdateTime           string `json:"updateTime,omitempty" gorm:"column:updateTime"`
}

func (Task) TableName() string {
	return "task"
}
