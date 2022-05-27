package database

import (
	"fmt"
	"log"
	"os"
	"server/pkg/ulti"
	"time"

	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbManipulation struct {
	DbManipulate
	DB *gorm.DB
}

var instance *dbManipulation

func GetConnectionInstance() *dbManipulation {
	DB := connection()
	if instance == nil {
		instance = &dbManipulation{DB: DB}
		return instance
	}
	return instance
}

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

func connection() (db *gorm.DB) {
	c := &Config{}
	result, err := ulti.ReadFileYaml("./configs/config_server.yaml")

	if err != nil {
		fmt.Println(err)
	} else {
		mapstructure.Decode(*result, c)
	}
	dbInfo := c.Config.Database.Environment


	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   logger.Silent, // Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  Colorful:                  false,          // Disable color
		},
	  )
	  
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbInfo.POSTGRES_HOST, dbInfo.Ports, dbInfo.POSTGRES_USER, dbInfo.POSTGRES_PASSWORD, dbInfo.POSTGRES_DB)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{Logger: newLogger,})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}

	return
}

