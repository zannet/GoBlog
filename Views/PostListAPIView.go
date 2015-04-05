package Views

import (
	//"net/http"
	"github.com/gin-gonic/gin"
)

func PostListAPIView(c *gin.Context) {
	// Get the needed data
	obj := c.MustGet("data").(gin.H)
	// And return it with the 'home' template
	c.JSON(200, obj)
}