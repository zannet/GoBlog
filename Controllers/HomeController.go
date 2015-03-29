package Controllers

import (
	//"html/template"
	"github.com/gin-gonic/gin"
	"../Models"
)

func HomeController(c *gin.Context) {
	// Create a new gin.H
	obj := gin.H{}
	// Set the web's title
	obj["title"] = "Sirikon's Lab"

	// Posts array to be filled
	var posts []Models.Post
	Models.Connection.Find(&posts)

	// Set the 'posts' in the data object
	obj["posts"] = posts
	// Store it in the context
	c.Set("data", obj)
}