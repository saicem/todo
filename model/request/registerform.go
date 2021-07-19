package request

type RegisterForm struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Captcha  string `json:"captcha,omitempty"`
}
