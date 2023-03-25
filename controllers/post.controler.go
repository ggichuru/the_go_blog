package controllers

import (
	"net/http"
	"strconv"
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

// [...] Update post handler
func (pc *PostController) UpdatePost(ctx *gin.Context) {
	// Extract post ID from the request parameters
	postId := ctx.Param("postId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdatePost
	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedPost models.Post

	// Check if the post ID exists
	result := pc.DB.First(&updatedPost, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	now := time.Now()

	// Construct a post argument
	postToUpdate := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		User:      currentUser.ID,
		CreatedAt: updatedPost.CreatedAt,
		UpdatedAt: now,
	}

	// Update post
	pc.DB.Model(&updatedPost).Updates(postToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPost})
}

// [...] Get a single post handler
func (pc *PostController) FindPostById(ctx *gin.Context) {
	// Get post ID from the params
	postId := ctx.Param("postId")

	var post models.Post
	result := pc.DB.First(&post, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "no post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
}

// [...] Get all posts handler
func (pc *PostController) FindPosts(ctx *gin.Context) {
	// configure for pagination
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var posts []models.Post

	// get all the post and apply pagination
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&posts)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": posts})
}

// [...] Delete a post
func (pc *PostController) DeletePost(ctx *gin.Context) {
	// Get id from param
	postId := ctx.Param("postId")

	// Delete post from DB
	result := pc.DB.Delete(&models.Post{}, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
