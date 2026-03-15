package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dbconfig "github.com/sidz111/jwt-twitter-msg/dbConfig"
	"github.com/sidz111/jwt-twitter-msg/models"
	"github.com/sidz111/jwt-twitter-msg/service"
	"github.com/sidz111/jwt-twitter-msg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	serv service.UserService
}

func NewAuthController(serv service.UserService) *AuthController {
	return &AuthController{serv: serv}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var user models.User
	var foundUser models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbconfig.DB.Model(&models.User{}).Where("id=?", user.ID).First(&foundUser)
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Password",
		})
		return
	}

	token, err := utils.GenerateJWT(user.Username, user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
