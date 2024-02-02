package controllers

import (
	"icomida/src/service"

	"github.com/gin-gonic/gin"
)

// @Summary Buscar um restaurante
// @Tags Restaurant
// @Description Encontra um restaurante
// @Accept json
// @Produce json
// @Param name query string true "Nome do Restaurante"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/v1/find-restaurant [get]
func FindRestaurants(c *gin.Context) {
	name := c.Query("name")

	filters := map[string]interface{}{
		"name": name,
	}
	service.FindRestaurants(c, filters)
}

// @Summary Criar novo restaurante
// @Tags Restaurant
// @Description Criação de novos restaurantes
// @Accept json
// @Produce json
// @Param name query string true "Nome do Restaurante"
// @Success 200 {string} json "{"message": "Restaurante criado"}"
// @Router /api/v1/create-restaurant [post]
func CreateRestaurants(c *gin.Context) {
	name := c.Query("Name")
	score := 0

	filters := map[string]interface{}{
		"name":  name,
		"score": score,
	}
	service.CreateRestaurant(c, filters)
}

// @Summary Buscar todos restaurante
// @Tags Restaurant
// @Description Encontra todos os restaurantes
// @Accept json
// @Produce json
// @Success 200 {string} json "{"message": "result"}"
// @Router /api/v1/all-restaurant [get]
func FindAllRestaurants(c *gin.Context) {
	service.AllRestaurants(c)
}
