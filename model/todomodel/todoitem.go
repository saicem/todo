package todomodel

import "time"

type TodoItem struct {
	TodoId      int       `json:"todo_id" gorm:"primaryKey"`
	UserId      int       `json:"user_id,omitempty"`
	TodoGroupId int       `json:"todo_group_id,omitempty"`
	TodoTitle   string    `json:"todo_title"`
	TodoContent string    `json:"todo_content,omitempty"`
	IsFinished  bool      `json:"is_finish,omitempty"`
	CreateAt    time.Time `json:"create_at" gorm:"autoCreateTime"`
}
