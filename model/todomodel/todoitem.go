package todomodel

import (
	"gorm.io/gorm"
	"time"
)

type TodoItem struct {
	gorm.Model
	Complete       bool
	Importance     int
	AnticipateTime time.Time
	Content        string
}
