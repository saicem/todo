package request

type LoginForm struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
