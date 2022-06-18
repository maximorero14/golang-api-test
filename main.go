package main

import (
	//"github.com/maximorero14/golang-api-test/models"
	"github.com/maximorero14/golang-api-test/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	//models.ConnectDatabase()

	// Routes
	r.GET("/books", services.FindBooks)
	r.GET("/books/:id", services.FindBook)
	r.POST("/books", services.CreateBook)
	r.PATCH("/books/:id", services.UpdateBook)
	r.DELETE("/books/:id", services.DeleteBook)

	// Run the server
	r.Run()
}