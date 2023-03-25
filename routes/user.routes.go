package routes

import (
	"github.com/ggichuru/the_go_blog/controllers"
	"github.com/ggichuru/the_go_blog/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (urc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("users")

	router.GET("/current_user", middleware.DeserializeUser(), urc.userController.GetCurrentUser)
}
