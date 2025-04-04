package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishavkumar7/restaurant-management/controllers"
)

func MenuRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/menus", controllers.GetMenus())
	incomingRouter.GET("/menus/menu-id", controllers.GetMenuWithId())
	incomingRouter.POST("/menu", controllers.CreateMenu())
	incomingRouter.PATCH("/menus/:menu-id", controllers.UpdateMenuWithId())
}
