package database

import (
	"gobackReactfront/go_auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:1234@/yt_go_auth"), &gorm.Config{})

	if err != nil {
		panic("could not connect to database")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
}
