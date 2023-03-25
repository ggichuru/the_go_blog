package routes

import (
	"github.com/ggichuru/the_go_blog/controllers"
	"github.com/ggichuru/the_go_blog/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (arc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/register", arc.authController.SignUpUser)
	router.POST("/login", arc.authController.SignInUser)
	router.POST("/refresh", arc.authController.RefreshAccesToken)
	router.GET("/logout", middleware.DeserializeUser(), arc.authController.LogoutUser)
}
