package response

import "github.com/saicem/todo/model/todomodel"

type TodoItemRes struct {
	ID         uint   `json:"id"`
	Complete   bool   `json:"complete,omitempty"`
	Importance int    `json:"importance,omitempty"`
	Content    string `json:"content,omitempty"`
	Tag        string `json:"tag,omitempty"`
}

func (r TodoItemRes) New(item *todomodel.TodoItem) *TodoItemRes {
	r.ID = item.ID
	r.Tag = item.Tag
	r.Complete = item.Complete
	r.Importance = item.Importance
	r.Content = item.Content
	return &r
}
