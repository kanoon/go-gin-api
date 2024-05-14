package controller

import (
	"demo-api/config"
	"demo-api/models"

	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context){
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
	// c.String(200, "Hello user!")
}