package register

import (
	"crypto/sha256"
	"fmt"
	"server/configs"
	"server/database"
	"server/pkg/utils"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)


func hasher(str string) string {
	byteStr := []byte(str)
	hashedStr := fmt.Sprintf("%x", sha256.Sum256(byteStr))
	return hashedStr
}
type userData struct {
	Name    string
	ID      int32
	Enabled bool
}
func Register(data database.User) (err error) {

	k := &configs.Keys{}
	result, e := utils.ReadFileYaml("./configs/key.yaml")

	if e != nil {
		fmt.Println(err)
	} else {
		mapstructure.Decode(*result, k)
	}

	data.Password = hasher(data.Password + k.Salt)

	mapData := structs.Map(data)

	db := database.GetConnectionInstance()
	err = db.InsertOneIntoTable("abc", "users", mapData)

	return
}