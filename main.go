package main

import (
	"assignment_2_golang/config"
	controller "assignment_2_golang/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.POST("/orders", controller.CreateOrder)
	r.GET("/orders", controller.GetAllOrders)
	r.PUT("/orders/:order_id", controller.UpdateOrder)
	r.DELETE("/orders/:order_id", controller.DeleteOrder)

	r.Run(":8000")
}
