package request

type TodoItemReq1 struct {
	TodoGroupId int    `json:"todo_group_id"`
	TodoTitle   string `json:"todo_title"`
	TodoContent string `json:"todo_content"`
}
type TodoItemReq2 struct {
	TodoGroupId int    `json:"todo_group_id"`
	TodoTitle   string `json:"todo_title"`
	TodoContent string `json:"todo_content"`
	IsFinished  bool   `json:"is_finished"`
}
