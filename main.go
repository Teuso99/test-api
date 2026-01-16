package main

import (
	"test-api/db"
	"test-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.GET("/items", handlers.GetItems)
	r.GET("/items/:id", handlers.GetItem)
	r.POST("/items", handlers.CreateItem)
	r.PUT("/items/:id", handlers.UpdateItem)
	r.DELETE("/items/:id", handlers.DeleteItem)

	r.Run(":8080")
}
