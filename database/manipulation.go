package database

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type DbManipulate interface {
	InsertOneIntoTable(schema string, tableName string, userData interface{}) (result *gorm.DB)
	GetColumnName(tableName string)(result *gorm.DB)
}


func (DbManipulation *dbManipulation) createDataString(data map[string]interface{}) (string, string){
	var columnNameStr string
	var dataStr string

	now := time.Now().Format("01-02-2006 15:04:05")

	for i, v := range data {

		if i == "Created_at"{
			columnNameStr += fmt.Sprintf(" %s,", i)
			dataStr += fmt.Sprintf(" '%v',", now)
		}else if i != "ID" && i != "Model"{
			dataStr += fmt.Sprintf(" '%v',", v)
			columnNameStr += fmt.Sprintf(" %s,", i)
		}
	}
	
	return strings.TrimSuffix(dataStr, ","), strings.TrimSuffix(columnNameStr, ",")
}
func (DbManipulation *dbManipulation) InsertOneIntoTable(schema string, tableName string, data map[string]interface{}) (err error){
	dataStr, columnNameStr := DbManipulation.createDataString(data)

	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, tableName, columnNameStr, dataStr)
	tx := DbManipulation.DB.Debug().Exec(query)
	err = tx.Error
	return
}

func (DbManipulation *dbManipulation) UsernameEmailIsExisted(schema string, tableName string, data User) bool {
	var result User
	query := fmt.Sprintf(`SELECT * FROM %s WHERE Username = '%s' OR Email = '%s'`, tableName, data.Username, data.Email)
	DbManipulation.DB.Raw(query).Scan(&result)
	if result.Username == ""{
		return false
	}
	return true
}

func (DbManipulation *dbManipulation) SelectUserFromTable(schema string, tableName string, data User) (result User){
	query := fmt.Sprintf(`SELECT * FROM %s WHERE Username = '%s' AND Password = '%s'`, tableName, data.Username, data.Password)
	DbManipulation.DB.Raw(query).Scan(&result)
	return
}

type result struct{
	ColumnName string
	DataType string
}

func (DbManipulation *dbManipulation) getColumnName(tableName string)(result []string){
	// result = DbManipulation.DB.Exec(getColumnNameQueryStr)
	query := fmt.Sprintf(`SELECT column_name FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = N'%s'`, tableName)
	DbManipulation.DB.Raw(query).Scan(&result)
	return
}
