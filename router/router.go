package router

import (
	"latihan-simple-rest-api-middleware/controllers"
	"latihan-simple-rest-api-middleware/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {

	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos", middlewares.Authentication())
	{
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/:photoId", controllers.GetOnePhoto)
		photoRouter.GET("/", controllers.GetPhotos)
		{
			photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
			photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
		}
	}

	commentRouter := r.Group("/comments", middlewares.Authentication())
	{
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/:commentId", controllers.GetOneComment)
		commentRouter.GET("/", controllers.GetComments)
		{
			commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
			commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
		}
	}

	socialMediaRouter := r.Group("/social-media", middlewares.Authentication())
	{
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/:socialMediaId", controllers.GetOneSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedias)
		{
			socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
			socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		}
	}

	return r

}
