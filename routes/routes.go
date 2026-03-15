package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/jwt-twitter-msg/controller"
	"github.com/sidz111/jwt-twitter-msg/middleware"
)

func SetRoutes(userController *controller.UserController, postController *controller.PostController, router *gin.Engine) *gin.Engine {
	r := gin.Default()
	user := r.Group("users")
	{
		user.POST("/", userController.CreateUser)
		user.GET("/:id", middleware.AuthMiddleware(), userController.GetUser)
		user.GET("/", middleware.AuthMiddleware(), userController.GetAllUsers)
		user.PUT("/", middleware.AuthMiddleware(), userController.UpdateUser)
		user.DELETE("/:id", middleware.AuthMiddleware(), userController.DeleteUser)
	}

	post := r.Group("posts")
	{
		post.POST("/", middleware.AuthMiddleware(), postController.CreatePost)
		post.GET("/:id", middleware.AuthMiddleware(), postController.GetPost)
		post.GET("/", middleware.AuthMiddleware(), postController.GetAllPosts)
		post.PUT("/", middleware.AuthMiddleware(), postController.UpdatePost)
		post.DELETE("/:id", middleware.AuthMiddleware(), postController.DeletePost)
		post.GET("/users/:userID", middleware.AuthMiddleware(), postController.GetPostsByUserId)
	}
	return r
}
