package main

import (
	//"github.com/maximorero14/golang-api-test/models"
	"github.com/gin-gonic/gin"
	"github.com/maximorero14/golang-api-test/src/services"
)

func main() {
	r := gin.Default()

	// Connect to database
	//models.ConnectDatabase()

	// Routes
	r.GET("/ping", services.Ping)
	r.GET("/activity", services.FindActivities)
	r.GET("/activity/:id", services.FindActivity)
	r.POST("/activity", services.CreateActivity)
	r.PATCH("/activity/:id", services.UpdateActivity)
	r.DELETE("/activity/:id", services.DeleteActivity)

	// Run the server
	r.Run(":8080")
}
