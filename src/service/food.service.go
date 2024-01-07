package service

import (
	"icomida/src/models"
	"icomida/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindFoods(c *gin.Context, filters map[string]interface{}) {
	result, err := repository.FindFood(models.DB, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})

}

func CreateFood(c *gin.Context, filters map[string]interface{}) {
	result, err := repository.AddFood(models.DB, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registro adicionado com sucesso"})
}
