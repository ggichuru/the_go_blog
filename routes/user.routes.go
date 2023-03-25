package routes

import (
	"github.com/ggichuru/the_go_blog/controllers"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (urc *UserRouteController) UserRoute(rg *gin.RouterGroup) {}
