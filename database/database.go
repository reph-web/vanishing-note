package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root@tcp(127.0.0.1:3306)/vanishing_note?charset=utf8mb4&parseTime=True&loc=Local" // Replace with YOUR MySQL connection string
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Cannot connect to database", err)
	}

	fmt.Println("✅ Connected to database")
}
