package routes

import (
	"golangOrderMatchingSystem/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine){
	api := r.Group("/api")
	{
		api.POST("/orders", controllers.PlaceOrder)
		// api.DELETE("/orders/:oderId" , controllers.CancelOrder)
		api.GET("/orderbook" , controllers.GetOrderBook)
		api.GET("/trades" , controllers.GetTradeHistory)
		api.GET("/orders/:orderId" , controllers.GetOrderStatus)
	}
}