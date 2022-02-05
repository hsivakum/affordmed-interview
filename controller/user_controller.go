package controller

import (
	"affordmed/models"
	"affordmed/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return userController{userService: userService}
}

func (u userController) CreateUser(ctx *gin.Context) {
	var signup models.User
	if err := ctx.ShouldBindJSON(&signup); err != nil {
		logrus.Errorf("unable to bind request body %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorCode": "ERR_INVALID_PAYLOAD",
		})
		return
	}

	err := u.userService.InsertUser(signup)
	if err != nil {
		logrus.Errorf("unable to create user %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorCode": "ERR_SOMETHING_WENT_WRONG",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

func (u userController) Login(ctx *gin.Context) {
	var login models.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		logrus.Errorf("unable to bind request body %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorCode": "ERR_INVALID_PAYLOAD",
		})
		return
	}

	err := u.userService.Login(login)
	if err != nil {
		logrus.Errorf("unable to login %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorCode": "ERR_SOMETHING_WENT_WRONG",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}
