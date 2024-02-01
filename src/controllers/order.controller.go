package controllers

import (
	"github.com/gin-gonic/gin"
)

// @Summary Obter todas as ordens
// @Tags Order
// @Description Obtém todas as ordens com base nos parâmetros de pesquisa
// @Accept json
// @Produce json
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/v1/all-order [get]
func getAllOrders(c *gin.Context) {

}

// @Summary Obter uma ordem
// @Tags Order
// @Description Obtém uma ordem com base nos parâmetros de pesquisa
// @Accept json
// @Produce json
// @Param orderNumber query string true "Numero do pedido"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/v1/find-order [get]
func getOneOrders(c *gin.Context) {

}

// @Summary Criar uma nova ordem
// @Tags Order
// @Description Cria uma nova ordem com base nos parâmetros fornecidos
// @Accept json
// @Produce json
// @Param RestaurantID query string true "Id do restaurante que criou a ordem de pedido pro usuario"
// @Param Items query array true "Comidas/Pedido"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/v1/create-order [post]
func createOrder(c *gin.Context) {

}

// @Summary Alterar o status de uma ordem
// @Tags Order
// @Description Altera o status de uma ordem com base nos parâmetros fornecidos
// @Accept json
// @Produce json
// @Param name query string true "Nome do alimento"
// @Param description query string false "Descrição do alimento"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/v1/status-order [put]
func changeStatusOrder(c *gin.Context) {

}
