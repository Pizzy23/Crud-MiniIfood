package controllers

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sua API de Comida
// @version 1.0
// @description API para gerenciar informações de alimentos
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api/v1
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/food", FindFood)
	r.POST("/food", CreateFood)

	r.GET("/all-restaurant", FindAllRestaurants)
	r.GET("/restaurant", FindRestaurants)
	r.POST("/restaurant", CreateRestaurants)

	return r
}
