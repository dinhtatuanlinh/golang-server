package models

type User struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}

type Response struct {
	Code    int
	Message []string
}

type SendingJson struct {
	Bar string
}