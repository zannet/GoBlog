package Views

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HomeView(c *gin.Context) {
	// Get the needed data
	obj := c.MustGet("data").(gin.H)
	// And return it with the 'home' template
	c.HTML(http.StatusOK, "home", obj)
}