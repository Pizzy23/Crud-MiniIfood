package main

import (
	"icomida/src/controllers"
	"icomida/src/models"
)

func main() {
	r := controllers.SetupRouter()

	models.ConnectDatabase()
	r.Run(":8080")
}
