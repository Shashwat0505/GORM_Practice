package main

import (
	"GORM_Project/router"
	"GORM_Project/services"
)

func init() {
	services.PostgresConnection()

}

func main() {
	router.Run("4000")
}
