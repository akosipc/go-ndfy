package main

import (
	"ndfy/config"
	"ndfy/models"
	"ndfy/routes"
)

func main() {
	db := config.SetupDatabase()
	db.AutoMigrate(&models.Artist{})

	r := routes.SetupRoutes(db)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
