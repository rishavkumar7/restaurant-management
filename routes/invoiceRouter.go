package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishavkumar7/restaurant-management/controllers"
)

func InvoiceRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/invoices", controllers.GetInvoices())
	incomingRouter.GET("/invoices/:invoice-id", controllers.GetInvoiceWithId())
	incomingRouter.POST("/invoice", controllers.CreateInvoice())
	incomingRouter.PATCH("/invoices/:invoice-id", controllers.UpdateInvoiceWithId())
}
