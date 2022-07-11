package models

type User struct {
	Username   string `json:"username", omitempty`
	Email      string `json:"email", omitempty`
	Password   string `json:"password"`
	Repassword string `json:"repassword", omitempty`
}

type Response struct {
	Code    int
	Message []string
}

type SendingJson struct {
	Bar string
}