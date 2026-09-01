package routes

import (
	"github.com/gin-gonic/gin"
	"Go-API/internal/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.BookById)
	router.POST("/books", handlers.CreateBook)
	router.PATCH("/checkout", handlers.CheckoutBook)
	router.PATCH("/return", handlers.ReturnBook)

	return router
}