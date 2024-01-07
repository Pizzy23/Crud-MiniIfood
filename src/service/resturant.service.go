package service

import (
	"icomida/src/models"
	"icomida/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindRestaurants(c *gin.Context, filters map[string]interface{}) {
	result, err := repository.FindRestaurants(models.DB, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})

}
func CreateRestaurant(c *gin.Context, filters map[string]interface{}) {
	result, err := repository.CreateRestaurant(models.DB, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Restaurante criado"})
}
func AllRestaurants(c *gin.Context) {
	result, err := repository.AllRestaurants(models.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})

}
