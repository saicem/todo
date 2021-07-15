package response

type Response struct {
	Status  ApiCode     `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
