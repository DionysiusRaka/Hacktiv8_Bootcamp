package routers

import (
	"Hactiv8_Bootcamp/H10-jwt/controllers"
	"Hactiv8_Bootcamp/H10-jwt/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Sukses tod",
			})
		})
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
	}

	return r
}
