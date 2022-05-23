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
	result, err := ulti.ReadFileYaml("./configs/key.yaml")

	if err != nil {
		fmt.Println(err)
	} else {
		mapstructure.Decode(*result, k)
	}
	

	db := database.GetConnectionInstance()

	data.Password = hasher(data.Password + k.Salt)
	r := db.InsertOneIntoTable("labs", "users", data)
	fmt.Println(r)
}