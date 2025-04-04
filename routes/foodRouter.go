package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishavkumar7/restaurant-management/controllers"
)

func FoodRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/foods", controllers.GetFoods())
	incomingRouter.GET("/foods/:food-id", controllers.GetFoodWithId())
	incomingRouter.POST("/food", controllers.CreateFood())
	incomingRouter.PATCH("/foods/:food-id", controllers.UpdateFoodWithId())
}
