package init

import "github.com/gin-gonic/gin"

func initRoutes(router *gin.Engine) {
	router.POST("/signup", userController.CreateUser)
	router.POST("/login", userController.Login)
}
