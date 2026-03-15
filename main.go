package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/jwt-twitter-msg/controller"
	dbconfig "github.com/sidz111/jwt-twitter-msg/dbConfig"
	"github.com/sidz111/jwt-twitter-msg/models"
	"github.com/sidz111/jwt-twitter-msg/repository"
	"github.com/sidz111/jwt-twitter-msg/routes"
	"github.com/sidz111/jwt-twitter-msg/service"
)

func main() {
	if err := dbconfig.ConnectDB(); err != nil {
		panic("failed to Connect DB")
	}

	r := gin.Default()
	dbconfig.DB.AutoMigrate(&models.User{}, &models.Post{})

	userRepo := repository.NewUserRepository(dbconfig.DB)
	userServ := service.NewUserService(userRepo)
	userController := controller.NewUserController(userServ)
	authController := controller.NewAuthController(userServ)

	postRepo := repository.NewPostRepository(dbconfig.DB)
	postServ := service.NewPostService(postRepo)
	postController := controller.NewPostController(postServ)

	route := routes.SetRoutes(userController, postController, authController, r)
	route.Run(":8080")

}
