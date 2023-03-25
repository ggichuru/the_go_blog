package routes

import (
	"github.com/ggichuru/the_go_blog/controllers"
	"github.com/ggichuru/the_go_blog/middleware"
	"github.com/gin-gonic/gin"
)

type PostRouteController struct {
	postController controllers.PostController
}

func NewPostRouteController(postController controllers.PostController) PostRouteController {
	return PostRouteController{postController}
}

func (prc *PostRouteController) PostRoute(rg *gin.RouterGroup) {
	router := rg.Group("posts")

	router.Use(middleware.DeserializeUser())

	router.POST("/", prc.postController.CreatePost)
	router.GET("/", prc.postController.FindPosts)
	router.PUT("/:postId", prc.postController.UpdatePost)
	router.GET("/:postId", prc.postController.FindPostById)
	router.DELETE("/:postId", prc.postController.DeletePost)
}
