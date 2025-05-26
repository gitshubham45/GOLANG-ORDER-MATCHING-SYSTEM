package main

import (
	"github.com/gin-gonic/gin"

	"golangOrderMatchingSystem/db"
)

func main() {

	db.InitDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.Run(":3000")
}
