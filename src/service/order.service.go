package service

import (
	"icomida/src/models"
	"icomida/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrder(c *gin.Context) {
	repository.FindOrder(models.DB)
}

func GetAllOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "WIP"})
}
func CreateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "WIP"})
}
func ChangeStatusOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "WIP"})
}
