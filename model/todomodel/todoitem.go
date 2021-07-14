package todomodel

import (
	"gorm.io/gorm"
)

type TodoItem struct {
	gorm.Model
	Complete   bool
	Importance int
	Content    string
	Uid        uint
}
