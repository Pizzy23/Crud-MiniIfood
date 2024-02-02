package controllers

import (
	"icomida/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sua API de Comida
// @version 1.0
// @description API para gerenciar informações de alimentos
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api
func SetupRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	v1 := r.Group("/api")
	{
		v1.GET("/food", FindFood)
		v1.POST("/food", CreateFood)

		v1.GET("/all-restaurant", FindAllRestaurants)
		v1.GET("/restaurant", FindRestaurants)
		v1.POST("/restaurant", CreateRestaurants)

		v1.GET("/all-order", getAllOrders)
		v1.GET("/find-order", getOneOrders)
		v1.POST("/create-order", createOrder)
		v1.PUT("/status-order", changeStatusOrder)

		v1.PUT("/login", connectUser)
		v1.PUT("/logged", disconnectUser)

		v1.POST("/create-user", createUser)
		v1.PUT("/edit-email", editEmail)
		v1.PUT("/edit-name", editName)
		v1.PUT("/edit-adress", editAdress)
		v1.GET("/get-user", getUser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
