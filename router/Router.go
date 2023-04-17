package router

import (
	"OctavianoRyan25/GO-JWT/controller"
	_ "OctavianoRyan25/GO-JWT/docs"
	"OctavianoRyan25/GO-JWT/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Social Media
// @version         1.0
// @description     This is a sample service for managing social media.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func StrartServer() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
	}
	
	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/create", controller.CreatePhoto)
		
		photoRouter.PUT("/update/:photoId", middleware.Authorization(), controller.UpdatePhoto)
		photoRouter.GET("/get", controller.GetAllPhoto)
		photoRouter.GET("/get/:photoId", middleware.Authorization(), controller.GetPhotobyID)
		photoRouter.DELETE("/delete/:photoId", middleware.Authorization(), controller.DeletePhoto)
	}
	
	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/create", controller.CreateComment)
		
		commentRouter.PUT("/update/:commentId", middleware.AuthorizationComment(), controller.UpdateComment)
		commentRouter.GET("/get", controller.GetAllComment)
		commentRouter.GET("/get/:commentId", middleware.AuthorizationComment(), controller.GetCommentbyID)
		commentRouter.DELETE("/delete/:commentId", middleware.AuthorizationComment(), controller.DeleteComment)
	}
	
	socialmediaRouter := r.Group("/socialmedia")
	{
		socialmediaRouter.Use(middleware.Authentication())
		socialmediaRouter.POST("/create", controller.CreateComment)
		
		socialmediaRouter.PUT("/update/:socialmediaId", middleware.AuthorizationComment(), controller.UpdateComment)
		socialmediaRouter.GET("/get", controller.GetAllComment)
		socialmediaRouter.GET("/get/:socialmediaId", middleware.AuthorizationComment(), controller.GetCommentbyID)
		socialmediaRouter.DELETE("/delete/:socialmediaId", middleware.AuthorizationComment(), controller.DeleteComment)
	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}