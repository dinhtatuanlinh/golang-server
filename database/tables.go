package database

var Tables = [string]map[string]map[string]string{
	"users": users,
}
var users = map[string]string{
	"id":         "SERIAL",
	"first_name": "text",
	"last_name":  "text",
	"username":   "text not null",
	"email":      "text not null",
	"password":   "text not null",
	"active":     "bool",
	"created_at": "date not null",
}