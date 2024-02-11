package routers

import (
	"Hactiv8_Bootcamp/H8-Assign/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	route := gin.Default()

	//Create
	route.POST("/createOrder/", controller.CreateOrder)
	route.GET("/getOrder/", controller.GetOrder)
	route.DELETE("/deleteOrder/:OrderId", controller.DeleteOrder)
	route.PUT("/updateOrder/:OrderId", controller.UpdateOrder)

	return route
}
