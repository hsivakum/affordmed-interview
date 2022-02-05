package init

import (
	"affordmed/controller"
	"affordmed/repository"
	"affordmed/service"
	"github.com/jmoiron/sqlx"
)

var userController controller.UserController

func initObjects(db *sqlx.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController = controller.NewUserController(userService)
}
