package routes

import "github.com/ggichuru/the_go_blog/controllers"

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}
