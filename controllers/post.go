package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/markokoen/go_api.git/models"
)

// GET /posts
func FindPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var posts []models.Post
	db.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// POST /posts
func CreatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input models.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create Post
	post := models.Post{Name: input.Name, Description: input.Description, ImageApp: input.ImageApp, ImageWeb: input.ImageWeb}
	db.Create(&post)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// GET /posts/:id
func FindPost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// PATCH /posts/:id
func UpdatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&post).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// DELETE /posts/:id
func DeletePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
