package routers

import (
	"Hactiv8_Bootcamp/H8-gin/controllers"
	_ "Hactiv8_Bootcamp/H8-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Car API
// @version 1.0
// @Description Simple sample service using swagger
// @termOfService http://swagger.io/terms/
// @contact.name NameJembut
// @contact.email email@swagger.io
// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html
//@host localhost:8080
//@BasePath /

func StartServer() *gin.Engine {
	router := gin.Default()
	//Create
	router.POST("/cars", controllers.CreateCar)
	//Update
	router.PATCH("/cars/:CarID", controllers.UpdateCar)
	//Get
	router.GET("/cars/getCars", controllers.GetCar)
	//Delete
	router.DELETE("/cars/:CarID", controllers.DeleteCar)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
