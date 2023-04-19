package services

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var err error

func PostgresConnection() {
	godotenv.Load(".env")

	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("DBUSERNAME")
	dbname := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, dbport)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("database connected successfully")
	}
	DB = db
}
