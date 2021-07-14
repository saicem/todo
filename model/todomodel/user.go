package todomodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NickName  string
	Password  string
	TodoItems []TodoItem `gorm:"foreignKey:ID"`
}
