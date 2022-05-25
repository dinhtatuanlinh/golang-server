package database

import (
	"fmt"
	"server/pkg/ulti"

	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type dbManipulation struct {
// 	DB *sql.DB
// }

// var instance *dbManipulation

// func GetConnectionInstance() *dbManipulation {
// 	DB := connection()
// 	if instance == nil {
// 		instance = &dbManipulation{DB: DB}
// 		CreateDatabase(instance, )
// 		return instance
// 	}
// 	return instance
// }

type Config struct {
	Config struct {
		Database struct {
			Environment struct {
				POSTGRES_HOST     string `yaml:"POSTGRES_HOST`
				POSTGRES_DB       string `yaml:"POSTGRES_DB"`
				POSTGRES_USER     string `yaml:"POSTGRES_USER"`
				POSTGRES_PASSWORD string `yaml:"POSTGRES_PASSWORD"`
				Ports             int64  `yaml:"ports"`
			} `yaml:"environment"`
		} `yaml:"database"`
	} `yaml:"config"`
}

func Connection() (db *gorm.DB) {
	c := &Config{}
	result, err := ulti.ReadFileYaml("./configs/config_server.yaml")

	if err != nil {
		fmt.Println(err)
	} else {
		mapstructure.Decode(*result, c)
	}
	dbInfo := c.Config.Database.Environment

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbInfo.POSTGRES_HOST, dbInfo.Ports, dbInfo.POSTGRES_USER, dbInfo.POSTGRES_PASSWORD, dbInfo.POSTGRES_DB)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	ok := db.Migrator().HasTable(&User{})
	if !ok{
		err := db.Migrator().CreateTable(&User{})
		if err != nil {
			fmt.Println(err)
		}
		// db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&User{})
		// ColumnTypes
	}
	
	return 
}
