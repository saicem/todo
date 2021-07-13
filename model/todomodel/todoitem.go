package todomodel

import "time"

type TodoItem struct {
	CreateTime     time.Time `json:"create_time"`
	Complete       bool      `json:"complete,omitempty"`
	Importance     int       `json:"importance,omitempty"`
	AnticipateTime time.Time `json:"anticipate_time"`
	Content        string    `json:"content,omitempty"`
}
