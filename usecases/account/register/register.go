package register

import (
	"crypto/sha256"
	"fmt"
	"server/database"
	"server/pkg/ulti"

	"github.com/mitchellh/mapstructure"
)
type keys struct {
	Secretkey string `yaml: "secretkey"`
	Salt string `yaml: "salt"`
}

func hasher(str string) string {
	byteStr := []byte(str)
	hashedStr := fmt.Sprintf("%x", sha256.Sum256(byteStr))
	return hashedStr
}
func Register(data *database.UserData) {

	k := &keys{}
	result, err := ulti.ReadFile("./configs/key.yaml")

	if err != nil {
		fmt.Println(err)
	} else {
		mapstructure.Decode(*result, k)
	}
	

	db := database.GetConnectionInstance()
	exist := db.CheckTableExist("labs", "users", db.DB)
	if !exist {
		fmt.Println("table is not existed")
	}
	data.Password = hasher(data.Password + k.Salt)
	fmt.Println(data.Password)
	r := db.InsertOneIntoTable("labs", "users", data, db.DB)
	fmt.Println(r)
}