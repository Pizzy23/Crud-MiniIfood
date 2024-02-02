package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "WIP"})
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
