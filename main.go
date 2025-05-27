package main

import (
	"github.com/gin-gonic/gin"

	"golangOrderMatchingSystem/controllers"
	"golangOrderMatchingSystem/db"
	"golangOrderMatchingSystem/routes"
)

func main() {

	db.InitDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	routes.OrderRoutes(r)

	port := controllers.GetPort()

	r.Run(":" + port)
}
