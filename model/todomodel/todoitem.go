package todomodel

import (
	"gorm.io/gorm"
)

type TodoItem struct {
	gorm.Model `json:"gorm_model"`
	Complete   bool   `json:"complete"`
	Importance int    `json:"importance"`
	Content    string `json:"content"`
	Tag        string `json:"tag"`
	Uid        uint   `json:"uid"`
}
