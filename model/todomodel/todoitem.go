package todomodel

import (
	"gorm.io/gorm"
)

type TodoItem struct {
	gorm.Model `json:"gorm_model"`
	Complete   bool   `json:"complete,omitempty"`
	Importance int    `json:"importance,omitempty"`
	Content    string `json:"content,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Uid        uint   `json:"uid,omitempty"`
}
