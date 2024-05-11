package main

import (
	"demo-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

	r.Use(cors.Default())

	
	routes.UserRoute(r)

	r.Run(":8001") 
}