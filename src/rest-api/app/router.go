package app

import (
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func mapURLs() {
	router.GET("/users", controllers.UsersController.GetAll)
	router.POST("/add", controllers.UsersController.Create)
	router.GET("/user/:id", controllers.UsersController.Get)
}

func StartApp() {
	mapURLs()
	router.Run(":8080")
}
