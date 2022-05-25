package database

import (
	"gorm.io/gorm"
)

var Tables = map[string]map[string]map[string]string{
	"userschema": userSchema,
}
var userSchema = map[string]map[string]string{
	"users": users,
}
var users = map[string]string{
	"id":         "SERIAL",
	"first_name": "text",
	"last_name":  "text",
	"username":   "text not null",
	"email":      "text not null",
	"password":   "text not null",
	"active":     "bool",
	"created_at": "date not null",
}

type User struct {
	gorm.Model
	ID         int `gorm:"primaryKey;autoIncrement"`
	First_name string
	Last_name  string
	Username   string `gorm:"not null"`
	Email      string `gorm:"not null"`
	Password   string `gorm:"not null"`
	ActivedAt string 
	CreatedAt string `gorm:"not null"`
	Status     string
}
// type ColumnType interface {
// 	Id() (serial int)
// 	First_name() string
// 	Last_name() string
// 	Name() string
// 	Username() string
// 	Email() string
// 	Password() string
// 	Active() bool
// 	Created_at() string
// 	Status() string

// 	DatabaseTypeName() string                 // varchar
// 	ColumnType() (columnType string, ok bool) // varchar(64)
// 	PrimaryKey() (isPrimaryKey bool, ok bool)
// 	AutoIncrement() (isAutoIncrement bool, ok bool)
// 	Length() (length int64, ok bool)
// 	DecimalSize() (precision int64, scale int64, ok bool)
// 	Nullable() (nullable bool, ok bool)
// 	Unique() (unique bool, ok bool)
// 	ScanType() reflect.Type
// 	Comment() (value string, ok bool)
// 	DefaultValue() (value string, ok bool)
// }
