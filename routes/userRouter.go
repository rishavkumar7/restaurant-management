package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishavkumar7/restaurant-management/controllers"
)

func UserRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/user", controllers.GetUsers())
	incomingRouter.GET("/user/:user-id", controllers.GetUserWithUserId())
	incomingRouter.POST("/user/signup", controllers.SignUp())
	incomingRouter.POST("/user/login", controllers.Login())
}
