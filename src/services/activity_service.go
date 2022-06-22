package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maximorero14/golang-api-test/src/internal"
	"github.com/maximorero14/golang-api-test/src/models"
)

// GET /books
// Find all books
func FindActivities(c *gin.Context) {
	var books []models.Activity
	//models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func Ping(c *gin.Context) {
	// Get model if exist

	/*if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}*/

	c.JSON(http.StatusOK, "pong")
}

// GET /books/:id
// Find a book
func FindActivity(c *gin.Context) {
	// Get model if exist
	activity, _ := internal.GetItem(c.Param("id"))
	/*if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}*/

	fmt.Printf("response activity.Id %s", activity.Id)

	c.JSON(http.StatusOK, gin.H{"ID": activity.Id, "Type": activity.Type, "UserId": activity.UserId})
}

// POST /books
// Create new book
func CreateActivity(c *gin.Context) {
	// Validate input
	/*var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Activity{Title: input.Title, Author: input.Author}
	//models.DB.Create(&book)*/

	c.JSON(http.StatusOK, gin.H{"data": nil})
}

// PATCH /books/:id
// Update a book
func UpdateActivity(c *gin.Context) {
	// Get model if exist
	/*var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}*/

	//models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": nil})
}

// DELETE /books/:id
// Delete a book
func DeleteActivity(c *gin.Context) {
	// Get model if exist
	//var book models.Book
	/*if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)*/

	c.JSON(http.StatusOK, gin.H{"data": true})
}
