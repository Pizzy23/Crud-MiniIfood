package service

import (
	"fmt"
	"icomida/src/models"
	"icomida/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context, filters map[string]interface{}) {
	password := filters["password"]
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.(string)), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Internal Error:", err)
		return
	}
	filters["password"] = hashedPassword
	result, err := repository.CreateUSer(models.DB, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result, "Response": "User Create"})
}
func EditName(c *gin.Context, filters map[string]interface{}) {
	err := repository.EditName(models.DB, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Name updated"})
}

func EditAdress(c *gin.Context, filters map[string]interface{}) {
	err := repository.EditAdress(models.DB, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Adress updated"})
}

func EditEmail(c *gin.Context, email string) {
	err := repository.EditEmail(models.DB, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "E-mail updated"})
}

func EditPassword(c *gin.Context, filters map[string]interface{}) {
	password := filters["password"]
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.(string)), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	filters["password"] = hashedPassword

	err = repository.EditPassword(models.DB, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated"})
}

func GetUser(c *gin.Context, email string) {
	result, err := repository.FindUserByEmail(models.DB, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}
