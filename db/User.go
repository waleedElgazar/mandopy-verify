package db

type User struct {
	Name  string `json:"name"`
	Otp   string `json:"otp"`
	Token string `json:"token"`
	Phone string `json:"phone"`
}


