package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishavkumar7/restaurant-management/controllers"
)

func OrderItemRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/order-items", controllers.GetOrderItems())
	incomingRouter.GET("/order-items/:order-item-id", controllers.GetOrderItemWithId())
	incomingRouter.GET("/:order-id/order-items", controllers.GetOrderItemsWithOrderId())
	incomingRouter.POST("/order-item", controllers.CreateOrderItem())
	incomingRouter.PATCH("/order-items/:order-item-id", controllers.UpdateOrderItemWithId())
}
