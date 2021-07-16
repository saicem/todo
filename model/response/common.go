package response

type Response struct {
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
