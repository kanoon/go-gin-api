package routes

import (
	"demo-api/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {

	r.GET("/users", controller.UserController)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}