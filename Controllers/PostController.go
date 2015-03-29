package Controllers

import (
	//"html/template"
	"github.com/gin-gonic/gin"
	"../Models"
)

func PostController(c *gin.Context) {
	// Create a new gin.H
	obj := gin.H{}
	// Set the web's title
	obj["title"] = "Sirikon's Lab"

	// Post element
	var post Models.Post
	Models.Connection.Where(&Models.Post{Slug: "prueba"}).First(&post)

	// Set the 'post' in the data object
	obj["post"] = post
	// Store it in the context
	c.Set("data", obj)
}