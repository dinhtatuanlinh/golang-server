package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"server/pkg/ulti"
// 	"time"

// 	_ "github.com/lib/pq"
// 	"github.com/mitchellh/mapstructure"
// )

// type DbManipulate interface {
// 	CreateTable(schema string, name string, db *sql.DB)
// 	CreateSchema(name string, db *sql.DB)
// 	CheckTableExist(schema string, name string, db *sql.DB) (exist bool)
// 	InsertOneIntoTable(schema string, tableName string, userData *UserData, db *sql.DB) (result bool)
// }

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

// func CreateDatabase(DBM *dbManipulation) {
// 	var interfaceSchemas []interface{}
// 	var requiredSchemas []string
// 	var requiredTables []string

// 	schemas := DBM.GetSchemas()
// 	for _, v := range schemas{
// 		interfaceSchemas = append(interfaceSchemas, v)
// 	}

// 	for i, _ := range Tables {
// 		requiredSchemas = append(requiredSchemas, i)
// 	}

// 	for _, v := range requiredSchemas{
// 		ok := ulti.CheckElementExistingInArray(v, interfaceSchemas)
// 		if !ok {
// 			DBM.CreateSchema(v)
// 		}
// 	}

// 	for _, v := range requiredSchemas{
// 		fmt.Println(Tables[v])
// 	}

// 	// for _, v := range schemas {
// 	// 	if schemas == v {
// 	// 		existedSchema = true
// 	// 	}
// 	// }

// 	// if existedSchema == false {
// 	// 	DBM.CreateSchema(schemas)
// 	// }

// 	// if DBM.CheckTableExist(schema, table) == false {
// 	// 	DBM.CreateTable(schema, table)
// 	// }

// }

// type Config struct {
// 	Config struct {
// 		Database struct {
// 			Environment struct {
// 				POSTGRES_HOST     string `yaml:"POSTGRES_HOST`
// 				POSTGRES_DB       string `yaml:"POSTGRES_DB"`
// 				POSTGRES_USER     string `yaml:"POSTGRES_USER"`
// 				POSTGRES_PASSWORD string `yaml:"POSTGRES_PASSWORD"`
// 				Ports             int64  `yaml:"ports"`
// 			} `yaml:"environment"`
// 		} `yaml:"database"`
// 	} `yaml:"config"`
// }

// func connection() *sql.DB {
// 	c := &Config{}
// 	result, err := ulti.ReadFileYaml("./configs/config_server.yaml")

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		mapstructure.Decode(*result, c)
// 	}
// 	dbInfo := c.Config.Database.Environment

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		dbInfo.POSTGRES_HOST, dbInfo.Ports, dbInfo.POSTGRES_USER, dbInfo.POSTGRES_PASSWORD, dbInfo.POSTGRES_DB)

// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }
// func (DbManipulation *dbManipulation) GetSchemas() []string {
// 	rows, err := DbManipulation.DB.Query("SELECT nspname FROM pg_catalog.pg_namespace")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	var results = make([]string, 0)
// 	for rows.Next() {
// 		var schema string
// 		if err := rows.Scan(&schema); err != nil {
// 			panic(err)
// 		}
// 		results = append(results, schema)
// 	}
// 	return results
// }

// func (DbManipulation *dbManipulation) CreateTable(schema string, name string) {
// 	count := 0
// 	last := len(Tables["users"])
// 	var tableFieldStr string
// 	for i, v := range Tables["users"] {
// 		count++
// 		if count < last {
// 			tableFieldStr += i + " " + v + ","
// 		} else {
// 			tableFieldStr += i + " " + v
// 		}

// 	}
// 	fmt.Println(tableFieldStr)
// 	query := fmt.Sprintf("CREATE TABLE %s.%s ("+
// 		"id SERIAL,"+
// 		"first_name text,"+
// 		"last_name text,"+
// 		"username text not null,"+
// 		"email text not null,"+
// 		"password text not null,"+
// 		"active bool,"+
// 		"created_at date not null)", schema, name)
// 	_, err := DbManipulation.DB.Exec(query)
// 	if err != nil {
// 		fmt.Println("creating table error")
// 	}
// 	fmt.Println("create users table successfully")
// }

// func (DbManipulation *dbManipulation) CreateSchema(name string) {
// 	query := fmt.Sprintf("CREATE SCHEMA %s", name)
// 	result, err := DbManipulation.DB.Exec(query)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// }

// func (DbManipulation *dbManipulation) CheckTableExist(schema string, name string) (exist bool) {
// 	query := fmt.Sprintf(
// 		"SELECT EXISTS ( SELECT * FROM %s.%s)", schema, name)
// 	_, err := DbManipulation.DB.Exec(query)
// 	if err != nil {
// 		exist = false
// 		return
// 	}
// 	exist = true
// 	return
// }

// type UserData struct {
// 	First_name string
// 	Last_name  string
// 	Username   string
// 	Email      string
// 	Password   string
// }

// func (DbManipulation *dbManipulation) InsertOneIntoTable(schema string, tableName string, userData *UserData) (result bool) {
// 	now := time.Now()
// 	dt := now.Format("01-02-2006 15:04:05")
// 	query := fmt.Sprintf("INSERT INTO %s.%s (first_name, last_name, username, email, password, active, created_at) "+
// 		"VALUES('%s', '%s', '%s', '%s', '%s', %v, '%s')", schema, tableName, userData.First_name, userData.Last_name, userData.Username,
// 		userData.Email, userData.Password, false, dt)
// 	_, err := DbManipulation.DB.Exec(query)
// 	if err != nil {
// 		result = false
// 		return
// 	}
// 	result = true
// 	return
// }
