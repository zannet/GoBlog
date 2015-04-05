package Controllers

import (
	//"html/template"
	"github.com/gin-gonic/gin"
	"../Models"
)

func PostListController(c *gin.Context) {
	// Create a new gin.H
	obj := gin.H{}

	// Posts array to be filled
	var posts []Models.Post
	Models.Connection.Find(&posts)

	// Set the 'posts' in the data object
	obj["posts"] = posts
	// Store it in the context
	c.Set("data", obj)
}