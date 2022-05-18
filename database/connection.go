package database

import (
	"database/sql"
	"fmt"
	"server/configs"
	"server/pkg/ulti"
	"time"

	_ "github.com/lib/pq"
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
	DB := connection()
	if instance == nil {
		instance = &dbManipulation{DB: DB}
	}
	return instance
}

func connection() *sql.DB {
	var c *configs.Config
	c, err := ulti.ReadFile("./configs/config_server.yaml")
	if err != nil {
		fmt.Println(err)
	}
	dbInfo := c.Config.Database.Environment

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		"localhost", dbInfo.Ports, dbInfo.POSTGRES_USER, dbInfo.POSTGRES_PASSWORD, dbInfo.POSTGRES_DB)

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
		exist = false
	}
	exist = true
	return
}

type UserData struct {
	first_name string
	last_name  string
	username   string
	email      string
	password   string
}

func (DbManipulation *dbManipulation) InsertOneIntoTable(schema string, tableName string, userData *UserData, db *sql.DB) (result bool) {
	dt := time.Now()
	query := fmt.Sprintf("INSERT INTO %s.%s (first_name, last_name, username, email, password, active, created_at) "+
		"VALUES($1, $2, $3, $4, $5, $6, $7)", schema, tableName)
	_, err := db.Exec(query, userData.first_name, userData.last_name, userData.username,
		userData.email, userData.password, false, dt.Format("01-02-2006 15:04:05"))
	if err != nil {
		result = false
	}
	result = true
	return
}
