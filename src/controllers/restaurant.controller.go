package controllers

import (
	"icomida/src/service"

	"github.com/gin-gonic/gin"
)

// @Summary Buscar um restaurante
// @Description Encontra alimentos com base nos parâmetros de pesquisa
// @Accept json
// @Produce json
// @Param name query string true "Nome do alimento"
// @Param description query string false "Descrição do alimento"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/v1/find-restaurant [get]
func FindRestaurants(c *gin.Context) {
	name := c.GetHeader("name")

	filters := map[string]interface{}{
		"name": name,
	}
	service.FindRestaurants(c, filters)
}

// @Summary Criar novo restaurante
// @Description Criação de novos restaurantes
// @Accept json
// @Produce json
// @Param name query string true "Nome de restaurante"
// @Success 200 {string} json "{"message": "Restaurante criado"}"
// @Router /api/v1/find-restaurant [get]
func CreateRestaurants(c *gin.Context) {
	name := c.GetHeader("Name")
	score := 0

	filters := map[string]interface{}{
		"name":  name,
		"score": score,
	}
	service.CreateRestaurant(c, filters)
}

// @Summary Buscar todos restaurante
// @Description Encontra alimentos com base nos parâmetros de pesquisa
// @Accept json
// @Produce json
// @Param name query string true "Nome do alimento"
// @Param description query string false "Descrição do alimento"
// @Success 200 {string} json "{"message": "result"}"
// @Router /api/v1/all-restaurant [get]
func FindAllRestaurants(c *gin.Context) {
	service.AllRestaurants(c)
}
