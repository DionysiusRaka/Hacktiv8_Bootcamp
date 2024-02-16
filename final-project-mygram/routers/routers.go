package routers

import (
	"Hactiv8_Bootcamp/final-project-mygram/controllers"
	"Hactiv8_Bootcamp/final-project-mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouters := r.Group("/users")
	{
		userRouters.POST("/register", controllers.UserRegister)
		userRouters.POST("/login", controllers.UserLogin)

	}

	photoRouters := r.Group("/photos")
	{
		photoRouters.Use(middlewares.Authentication())
		photoRouters.POST("/", controllers.CreatePhoto)
		photoRouters.GET("/", controllers.GetAllPhotos)
		photoRouters.GET("/:photoId", controllers.GetPhotoByID)
		photoRouters.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouters.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouters := r.Group("/comments")
	{
		commentRouters.Use(middlewares.Authentication())
		commentRouters.GET("/", controllers.GetAllComments)
		commentRouters.GET("/:commentId", controllers.GetCommentByID)
		commentRouters.POST("/:photoId", controllers.CreateComment)
		commentRouters.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouters.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socMedRouters := r.Group("/socmed")
	{
		socMedRouters.Use(middlewares.Authentication())
		socMedRouters.POST("/", controllers.CreateSocMed)
		socMedRouters.GET("/", controllers.GetAllSocMed)
		socMedRouters.GET("/:socMedId", controllers.GetSocMedByID)
		socMedRouters.PUT("/:socMedId", middlewares.SocMedAuthorization(), controllers.UpdateSocMed)
		socMedRouters.DELETE("/:socMedId", middlewares.SocMedAuthorization(), controllers.DeleteSocMed)
	}

	return r
}
