package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/markokoen/go_api.git/controllers"
	"github.com/markokoen/go_api.git/models"
)

// InitRoutes initialising routes
func InitRoutes() {

	router := gin.Default()
	db := models.SetupModels() // new

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	api := router.Group("api")

	PostRoutes(api)
	router.Run()
}

// PostRoutes routes for posts
func PostRoutes(api *gin.RouterGroup) {
	r := api.Group("posts")
	r.GET("", controllers.FindPosts)
	r.GET(":id", controllers.FindPost)
	r.POST("", controllers.CreatePost)
	r.PUT(":id", controllers.UpdatePost)
	r.DELETE(":id", controllers.DeletePost)
}
