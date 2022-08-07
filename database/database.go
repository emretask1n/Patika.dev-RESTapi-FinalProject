package database

import (
	"REST_API/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func ConnectDb() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:1234@tcp(127.0.0.1:3306)/restapi?charset=utf8mb4&parseTime=True&loc=Local"
	/*
		NOTE:
		To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
		To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	*/

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	// db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.ShoppingCart{})
	Instance = db

}
