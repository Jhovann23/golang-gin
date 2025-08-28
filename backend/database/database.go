package database

import (
	"backend/config"
	"backend/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	//setup konfig env/load konfig env
	dbUser := config.GetEnv("DB_USER", "root")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPass := config.GetEnv("DB_PASS", "")
	dbPort := config.GetEnv("DB_PORT", "3306")
	dbName := config.GetEnv("DB_NAME", "db_golang")

	//data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	dbUser, dbPass, dbHost, dbPort, dbName)

	//connect DB
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to database", err)
	}
	fmt.Println("Database connected successfully!")

	//migrate DB
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Fatal to migrate database:", err)
	}
	fmt.Println("Successfully to migrate database!")
}