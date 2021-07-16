package request

type TodoItemReq struct {
	Importance int    `json:"importance"`
	Content    string `json:"content"`
	Tag        string `json:"tag"`
}
