package Controllers

import (
	"github.com/gin-gonic/gin"
	"../Models"
)

type PostJSON struct {
    Title string
    Body  string
}

func PostUpdateController(c *gin.Context) {
	// Create a new gin.H
	obj := gin.H{}

	var json PostJSON
	c.Bind(&json)

	// Post element
	var post Models.Post
	Models.Connection.Where(&Models.Post{Slug: c.Params.ByName("slug")}).First(&post)

	post.Title = json.Title
	post.Body = json.Body

	Models.Connection.Save(&post)

	// Set the 'post' in the data object
	obj["post"] = post
	// Store it in the context
	c.Set("data", obj)
}