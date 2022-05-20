package database

import (
	"database/sql"
	"fmt"
	"server/pkg/ulti"
	"time"

	_ "github.com/lib/pq"
	"github.com/mitchellh/mapstructure"
)

type DbManipulate interface {
	CreateTable(schema string, name string, db *sql.DB)
	CreateSchema(name string, db *sql.DB)
	CheckTableExist(schema string, name string, db *sql.DB) (exist bool)
	InsertOneIntoTable(schema string, tableName string, userData *UserData, db *sql.DB) (result bool)
}

type dbManipulation struct {
	DB *sql.DB
}

var instance *dbManipulation

func GetConnectionInstance() *dbManipulation {
	DB := Connection()
	if instance == nil {
		instance = &dbManipulation{DB: DB}
	}
	return instance
}

type Config struct {
	Config struct {
		Database struct {
			Environment struct {
				POSTGRES_HOST string `yaml:"POSTGRES_HOST`
				POSTGRES_DB       string `yaml:"POSTGRES_DB"`
				POSTGRES_USER     string `yaml:"POSTGRES_USER"`
				POSTGRES_PASSWORD string  `yaml:"POSTGRES_PASSWORD"`
				Ports             int64  `yaml:"ports"`
			} `yaml:"environment"`
		} `yaml:"database"`
	} `yaml:"config"`
}

func Connection() *sql.DB {

	c := &Config{}
	result, err := ulti.ReadFile("./configs/config_server.yaml")

	if err != nil {
		fmt.Println(err)
	} else {
		mapstructure.Decode(*result, c)
	}
	dbInfo := c.Config.Database.Environment


	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	dbInfo.POSTGRES_HOST, dbInfo.Ports, dbInfo.POSTGRES_USER, dbInfo.POSTGRES_PASSWORD, dbInfo.POSTGRES_DB)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func (DbManipulation *dbManipulation) CreateTable(schema string, name string, db *sql.DB) {
	query := fmt.Sprintf("CREATE TABLE %s.%s ("+
		"id SERIAL,"+
		"first_name text,"+
		"last_name text,"+
		"username text not null,"+
		"email text not null,"+
		"password text not null,"+
		"active bool,"+
		"created_at date not null)", schema, name)
	result, err := db.Exec(query)
	if err != nil {
		fmt.Println("creating table error")
	}
	fmt.Println(result)
}

func (DbManipulation *dbManipulation) CreateSchema(name string, db *sql.DB) {
	query := fmt.Sprintf("CREATE SCHEMA %s", name)
	result, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func (DbManipulation *dbManipulation) CheckTableExist(schema string, name string, db *sql.DB) (exist bool) {
	query := fmt.Sprintf(
		"SELECT EXISTS ( SELECT * FROM %s.%s)", schema, name)
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		exist = false
		return
	}
	exist = true
	return
}

type UserData struct {
	First_name string
	Last_name  string
	Username   string
	Email      string
	Password   string
}

func (DbManipulation *dbManipulation) InsertOneIntoTable(schema string, tableName string, userData *UserData, db *sql.DB) (result bool) {
	now := time.Now()
	dt := now.Format("01-02-2006 15:04:05")
	query := fmt.Sprintf("INSERT INTO %s.%s (first_name, last_name, username, email, password, active, created_at) "+
		"VALUES('%s', '%s', '%s', '%s', '%s', %v, '%s')", schema, tableName, userData.First_name, userData.Last_name, userData.Username,
		userData.Email, userData.Password, false, dt)
	_, err := db.Exec(query)
	if err != nil {
		result = false
		return
	}
	result = true
	return
}
