package request

type TodoItemReq struct {
	Complete   bool   `json:"complete,omitempty"`
	Importance int    `json:"importance,omitempty"`
	Content    string `json:"content,omitempty"`
	Uid        uint   `json:"uid,omitempty"`
}
