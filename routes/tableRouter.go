package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishavkumar7/restaurant-management/controllers"
)

func TableRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/tables", controllers.GetTables())
	incomingRouter.GET("/tables/:table-id", controllers.GetTableWithId())
	incomingRouter.POST("/table", controllers.CreateTable())
	incomingRouter.PATCH("/tables/:table-id", controllers.UpdateTableWithId())
}
