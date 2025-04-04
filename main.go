package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rishavkumar7/restaurant-management/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" { // setting default port if port is not specified
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	router.Run(":" + port)
}
