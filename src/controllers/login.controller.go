package controllers

import (
	"icomida/src/service"

	"github.com/gin-gonic/gin"
)

// @Summary Conectar o usuario
// @Tags Login
// @Description Altera o modo de login para true
// @Accept json
// @Produce json
// @Param Email query string true "Email para pegar o atributo do login"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/login [put]
func connectUser(c *gin.Context) {
	email := c.Query("email")
	service.ConnectByEmail(c, email)
}

// @Summary Desconectar o usuario
// @Tags Login
// @Description Altera o modo de login para false
// @Accept json
// @Produce json
// @Param Email query string true "Email para pegar o atributo do login"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/loggout [put]
func disconnectUser(c *gin.Context) {
	email := c.Query("email")
	service.DisconnectUser(c, email)
}
