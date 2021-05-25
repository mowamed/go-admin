package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Connect() {
	_, err := gorm.Open(mysql.Open(os.Getenv("GO_ADMIN_DB")), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
}
