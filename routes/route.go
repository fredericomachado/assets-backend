package routes

import (
	"assets-backend/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/hello", controllers.Hello)
	r.POST("/categories", controllers.Save)
	r.Run()
}
