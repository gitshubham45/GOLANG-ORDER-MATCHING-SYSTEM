package routes

import (
	"golangOrderMatchingSystem/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine){
	api := r.Group("/api")
	{
		api.GET("/orders", controllers.GetOrders)
		api.POST("/orders", controllers.PlaceOrder)
		api.DELETE("/orders/:orderId" , controllers.CancelOrder)
		api.GET("/orderbook" , controllers.GetOrderBook)

		api.GET("/trades" , controllers.GetTradeHistory)
		api.GET("/orders/:orderId" , controllers.GetOrderStatus)
	}
}