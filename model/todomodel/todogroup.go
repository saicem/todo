package todomodel

import "time"

type TodoGroup struct {
	TodoGroupId   int        `json:"todo_group_id,omitempty" gorm:"primaryKey"`
	TodoGroupName string     `json:"todo_group_name,omitempty"`
	UserId        int        `json:"user_id"`
	TodoItems     []TodoItem `json:"todo_items,omitempty" gorm:"foreignKey:TodoGroupId;references:TodoGroupId"`
	CreateAt      time.Time  `json:"create_at" gorm:"autoCreateTime"`
}
