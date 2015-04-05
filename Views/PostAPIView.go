package Views

import (
	"github.com/gin-gonic/gin"
)

func PostAPIView(c *gin.Context) {
	// Get the needed data
	obj := c.MustGet("data").(gin.H)
	// And return it with the 'home' template
	c.JSON(200, obj)
}