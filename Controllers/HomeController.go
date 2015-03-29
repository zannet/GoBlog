package Controllers

import (
	//"html/template"
	"github.com/gin-gonic/gin"
	"../Models"
)

func HomeController(c *gin.Context) {
	obj := gin.H{}
	obj["title"] = "Sirikon's Lab"

	var posts []Models.Post
	Models.Connection.Find(&posts)

	obj["posts"] = posts
	c.Set("data", obj)
}