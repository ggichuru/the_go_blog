package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/ggichuru/the_go_blog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Create a struct to extend gorm.DB -> this allows us to have access to all APIs the gorm lib exposes
type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

// [...] create post handler
func (pc *PostController) CreatePost(ctx *gin.Context) {
	// Get the authenticated user's credentials from the context object
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreatePostRequest

	// Validate the request body before defining the post struct args
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newPost := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		User:      currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// insert the new record to DB
	result := pc.DB.Create(&newPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	// Return a JSON response with the newly created post to the client
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPost})
}
