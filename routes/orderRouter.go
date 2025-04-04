package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishavkumar7/restaurant-management/controllers"
)

func OrderRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/orders", controllers.GetOrders())
	incomingRouter.GET("/orders/order-id", controllers.GetOrderWithId())
	incomingRouter.POST("/order", controllers.CreateOrder())
	incomingRouter.PATCH("/orders/:order-id", controllers.UpdateOrderWithId())
}
