package main

import (
	"log"

	"Go-API/internal/routes"
)

func main() {
	router := routes.SetupRouter()

	log.Println("Server running on http://localhost:8080")
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}