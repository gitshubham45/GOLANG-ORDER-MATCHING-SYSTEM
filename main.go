package main

import (  
	"fmt" 
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("URL Shortner")
	r := gin.Default()

	r.GET("/" , func(c *gin.Context){
		c.JSON(200 , gin.H{
			"message" : "ok",
		})
	})

	r.Run(":3000")
}