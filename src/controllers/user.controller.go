package controllers

import (
	"github.com/gin-gonic/gin"
)

// @Summary Criar um novo usuario
// @Tags User
// @Description Criar um usuario a partir das novas infos
// @Accept json
// @Produce json
// @Param Email query string true "Email do usuario"
// @Param Name query string true "Nome do usuario"
// @Param Adress query string true "Endereço do usuario"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/create-user [post]
func createUser(c *gin.Context) {

}

// @Summary Editar usuario (Email)
// @Tags User
// @Description Editar um usuario a partir das novas infos
// @Accept json
// @Produce json
// @Param Email query string false "Email para editar as novas info"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/edit-email [put]
func editEmail(c *gin.Context) {

}

// @Summary Editar usuario (Name)
// @Tags User
// @Description Editar um usuario a partir das novas infos
// @Accept json
// @Produce json
// @Param Name query string false "Name para editar as novas info"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/edit-name [put]
func editName(c *gin.Context) {

}

// @Summary Editar usuario (Adress)
// @Tags User
// @Description Editar um usuario a partir das novas infos
// @Accept json
// @Produce json
// @Param Adress query string false "Adress para editar as novas info"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/edit-adress [put]
func editAdress(c *gin.Context) {

}

// @Summary Pegar um usuario
// @Tags User
// @Description Buscar um usuario dentro da db
// @Accept json
// @Produce json
// @Param Email query string true "Email para pegar um usuario especifico"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/get-user [get]
func getUser(c *gin.Context) {

}
