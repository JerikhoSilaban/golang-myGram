package routers

import (
	"DTSGolang/FinalProject/controllers"
	"DTSGolang/FinalProject/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "DTSGolang/FinalProject/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

// @title Mygram (Instagram Clone)
// @version 1.0
// @description This is an Instagram Clone for CRUD-ing photos, and comments among users
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8000
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "MyGram (Instagram Clone) by jerikhos")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	guestRouter := r.Group("/users")
	{
		// Register user
		guestRouter.POST("/register", controllers.UserRegister)

		// Login user
		guestRouter.POST("/login", controllers.UserLogin)
	}

	r.Use(middlewares.Authentication())
	{

		socialMediaRouter := r.Group("/socialMedia")
		{
			// Post social media
			socialMediaRouter.POST("/", controllers.CreateSocialMedia)

			// Get social media by id
			socialMediaRouter.GET("/:socialMediaID", controllers.GetSocialMedia)

			// Get all social media
			socialMediaRouter.GET("/", controllers.GetSocialMedias)

			// Get all social media corresponding user
			socialMediaRouter.GET("/my", controllers.GetMySocialMedias)

			// Update social media by id
			socialMediaRouter.PUT("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)

			// delete social media by id
			socialMediaRouter.DELETE("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
		}

		photoRouter := r.Group("/photos")
		{
			// Post photo
			photoRouter.POST("/", controllers.CreatePhoto)

			// Get photo by id
			photoRouter.GET("/:photoID", controllers.GetPhoto)

			// Get all photo corresponding user
			photoRouter.GET("/my", controllers.GetMyPhotos)

			// Get all photo
			photoRouter.GET("/", controllers.GetPhotos)

			// Update photo by id
			photoRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)

			// Delete photo by id
			photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
		}

		commentRouter := r.Group("/comments")
		{
			// Post comment
			commentRouter.POST("/", controllers.CreateComment)

			// Get comment by id
			commentRouter.GET("/:commentID", controllers.GetComment)

			// Get all comments corresponding user
			commentRouter.GET("/my", controllers.GetMyComments)

			// Get all comments
			commentRouter.GET("/", controllers.GetComments)

			// Update comment by id
			commentRouter.PUT("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)

			// Delete photo by id
			commentRouter.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
		}
	}

	return r
}
