package controllers

import (
	"icomida/src/service"

	"github.com/gin-gonic/gin"
)

// @Summary Encontrar alimentos
// @Tags Food
// @Description Encontra alimentos com base nos parâmetros de pesquisa
// @Accept json
// @Produce json
// @Param name query string true "Nome do alimento"
// @Param description query string false "Descrição do alimento"
// @Success 200 {string} json "{"message": "Alimentos encontrados"}"
// @Router /api/v1/food [get]
func FindFood(c *gin.Context) {
	name := c.Query("name")

	filters := map[string]interface{}{
		"name": name,
	}
	service.FindFoods(c, filters)
}

// @Summary Criar alimentos
// @Tags Food
// @Description Criar novos alimentos :)
// @Accept json
// @Produce json
// @Param name query string true "Nome do alimento"
// @Param description query string false "Descrição do alimento"
// @Param price query number true "Valor do Produto"
// @Success 200 {string} json "{"message": "Registro adicionado com sucesso"}"
// @Router /api/v1/food [post]
func CreateFood(c *gin.Context) {
	name := c.Query("Name")
	description := c.Query("Description")
	price := c.Query("Price")
	restaurant := c.Query("Restaurant")

	filters := map[string]interface{}{
		"name":        name,
		"description": description,
		"price":       price,
		"restaurant":  restaurant,
	}
	service.CreateFood(c, filters)
}
