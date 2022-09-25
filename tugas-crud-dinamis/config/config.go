package config

import (
	"day4/tugas-crud-dinamis/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// init connect to mysql db
func InitDB() {

	e := godotenv.Load("lokal.env")
	if e != nil {
		log.Fatalf("Erorr env. Err: %s", e)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		os.Getenv("DB_USER"),
		// "root",
		// os.Getenv("DB_PASSWORD"),
		"123456",
		// os.Getenv("DB_HOST"),
		"localhost",
		// os.Getenv("DB_PORT"),
		"3306",
		// os.Getenv("DB_NAME"),
		"agmc",
		"utf8mb4",
		"True",
		"Local",
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Success!")
}

// init function execute before main function
func init() {
	InitDB()
	InitialMigration()
}

// funtion for auto create table base on define struct
func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate((&models.Book{}))
}
