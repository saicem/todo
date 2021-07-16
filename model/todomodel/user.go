package todomodel

import "time"

type User struct {
	UserId     int    `gorm:"primaryKey"`
	Email      string `gorm:"unique"`
	NickName   string
	Password   string
	TodoGroups []TodoGroup `gorm:"foreignKey:UserId;references:UserId"`
	TodoItems  []TodoItem  `gorm:"foreignKey:UserId;references:UserId"`
	CreateAt   time.Time   `json:"create_at" gorm:"autoCreateTime"`
}
