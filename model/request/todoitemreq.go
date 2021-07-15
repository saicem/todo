package request

type TodoItemReq struct {
	Importance int    `json:"importance,omitempty"`
	Content    string `json:"content,omitempty"`
}
