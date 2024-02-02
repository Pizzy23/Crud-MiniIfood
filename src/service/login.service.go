package service

import (
	"icomida/src/models"
	"icomida/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConnectByEmail(c *gin.Context, email string) {
	result, err := repository.ConnectUserInDB(models.DB, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}

func DisconnectUser(c *gin.Context, email string) {
	result, err := repository.DesconnectUserInDB(models.DB, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}
