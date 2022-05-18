package register

import (
	"crypto/sha256"
	"fmt"
	"server/database"
)
type configs struct{
	config struct{
		salt string `yaml:"salt"`
	} `yaml:"config"`
}

func hasher(str string) string {
	byteStr := []byte(str)
	hashedStr := fmt.Sprintf("%x", sha256.Sum256(byteStr))
	return hashedStr
}
func Register(data *database.UserData) {
	db := database.GetConnectionInstance()
	_ = db.CheckTableExist("labs", "users", db.DB)

	data.password = hasher(data.password)
	db.InsertOneIntoTable("labs", "users", data, db.DB)
}