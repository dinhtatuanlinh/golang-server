package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         int `gorm:"PRIMARY_KEY AUTO_INCREMENT"`
	First_name string
	Last_name  string
	Username   string `gorm:"type: varchar(128); NOT NULL UNIQUE; default:null"`
	Email      string `gorm:"type: varchar(128); NOT NULL UNIQUE; default:null"`
	Password   string `gorm:"type: varchar(128); NOT NULL; default:null"`
	Avatar_url string
	Actived_at  string
	Created_at  string `gorm:"type: varchar(128); NOT NULL; default:null"`
	Status     string
	Delete_status string
}

