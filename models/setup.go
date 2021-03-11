package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", os.Getenv("MYSQL_CONNECTION_URL"))

	if err != nil {
		fmt.Println("Error: ", err)
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Task{})

	DB = database
}
