package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic("Failed to connect to database!")
	}
}
