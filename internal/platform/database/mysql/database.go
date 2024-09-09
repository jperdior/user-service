package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:verysecretpass@tcp(database:3306)/user-service?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	fmt.Println("Database connected")
	DB = db
}
