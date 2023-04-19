package main

import (
	"GORM_Project/models"
	"GORM_Project/services"
	"fmt"
	"log"
)

func init() {
	services.PostgresConnection()

}
func main() {
	//deletedErr := services.DB.Migrator().DropTable(&models.User{})
	migrationError := services.DB.AutoMigrate(&models.User{})
	if migrationError != nil {
		log.Fatal(migrationError)
	} else {
		fmt.Println("Migration successful")
	}

}
