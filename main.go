package main

import (
	"github.com/gin-gonic/gin"
	"github.com/markokoen/go_api.git/controllers"
	"github.com/markokoen/go_api.git/models"
)

func main() {
	r := gin.Default()
	db := models.SetupModels()

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/posts", controllers.FindPosts)
	r.POST("/posts", controllers.CreatePost)       // create
	r.GET("/posts/:id", controllers.FindPost)      // find by id
	r.PATCH("/posts/:id", controllers.UpdatePost)  // update by id
	r.DELETE("/posts/:id", controllers.DeletePost) // delete by id
	r.Run()
}
