package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Username  string `json:"username"`
	Content string `json:"content"`
	PostId int `json:"post_id"`
	Ip string `json:"ip"`
}

func (m *Comment) TableName() string {
	return TableName("comment")
}