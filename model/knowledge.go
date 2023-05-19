package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Tag []string

type Knowledge struct {
	Id         int    `gorm:"primary_key;column:id"  json:"id"`
	Title      string `gorm:"column:title"    json:"title"`
	Content    string `gorm:"column:content" json:"content"`
	CKey       string `gorm:"column:ckey" json:"key"`
	Hit        int    `gorm:"hit" json:"hit"`
	Shared     *bool  `gorm:"column:shared" json:"shared"`
	Author     string `gorm:"column:author" json:"author"`
	CreateTime string `gorm:"column:createTime" json:"createTime"`
	UpdateTime string `gorm:"column:updateTime" json:"updateTime"`
	Tags       Tag    `gorm:"column:tags" json:"tags,omitempty"`
	ParentId   int    `gorm:"column:parentId" json:"parentId"`
}

func (d *Knowledge) TableName() string {
	return "knowledge"
}

func (t *Tag) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, &t)
}

func (t Tag) Value() (driver.Value, error) {
	str, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return string(str), nil
}

type Category struct {
	Id       int    `gorm:"primary_key;column:id" json:"id"`
	ParentId int    `gorm:"column:parentId" json:"parentId"`
	Title    string `gorm:"column:title" json:"title"`
	IsLeaf   *bool  `gorm:"column:isLeaf" json:"isLeaf"`
	CKey     string `gorm:"column:ckey" json:"key"`
}

type CategoryTree struct {
	Id       int             `json:"id"`
	ParentId int             `json:"parentId"`
	Title    string          `json:"title,omitempty"`
	IsLeaf   *bool           `json:"isLeaf,omitempty"`
	Key      string          `json:"key,omitempty"`
	Children []*CategoryTree `json:"children,omitempty"`
}

func (d *Category) TableName() string {
	return "category"
}

type Knows struct {
	Href       string `json:"href"`
	Title      string `json:"title"`
	CreateTime string `json:"createTime"`
}

type ListData struct {
	Year  string  `json:"year"`
	Knows []Knows `json:"knows"`
}
