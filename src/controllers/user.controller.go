package controllers

import (
	"icomida/src/service"

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
	name := c.Query("Name")
	email := c.Query("Email")
	adress := c.Query("Adress")
	password := c.Query("Password")

	filters := map[string]interface{}{
		"name":     name,
		"email":    email,
		"adress":   adress,
		"password": password,
	}

	service.CreateUser(c, filters)
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
	email := c.Query("Email")

	service.EditEmail(c, email)
}

// @Summary Editar usuario (Name)
// @Tags User
// @Description Editar um usuario a partir das novas infos
// @Accept json
// @Produce json
// @Param Name query string false "Name para editar as novas info"
// @Param Email query string false "Email para editar as novas info"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/edit-name [put]
func editName(c *gin.Context) {
	name := c.Query("Name")
	email := c.Query("Email")
	filters := map[string]interface{}{
		"name":  name,
		"email": email,
	}

	service.EditName(c, filters)
}

// @Summary Editar usuario (Adress)
// @Tags User
// @Description Editar um usuario a partir das novas infos
// @Accept json
// @Produce json
// @Param Adress query string false "Adress para editar as novas info"
// @Param Email query string false "Email para editar as novas info"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/edit-adress [put]
func editAdress(c *gin.Context) {
	adress := c.Query("Adress")
	email := c.Query("Email")
	filters := map[string]interface{}{
		"adress": adress,
		"email":  email,
	}
	service.EditAdress(c, filters)
}

// @Summary Editar usuario (Password)
// @Tags User
// @Description Editar um usuario a partir das novas infos
// @Accept json
// @Produce json
// @Param Password query string false "Password para editar as novas info"
// @Param Email query string false "Email para editar as novas info"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/edit-adress [put]
func editPassword(c *gin.Context) {
	password := c.Query("Password")
	email := c.Query("Email")
	filters := map[string]interface{}{
		"password": password,
		"email":    email,
	}
	service.EditPassword(c, filters)
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
	email := c.Query("Email")

	service.GetUser(c, email)
}
