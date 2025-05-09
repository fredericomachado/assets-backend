package controllers

import (
	"assets-backend/database"
	"assets-backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Save(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not read body"})
		return
	}
	fmt.Println("Raw payload:", string(body))

	fmt.Print("category: ", category)

	database.DB.Create(&category)
	c.JSON(http.StatusOK, category)
}
