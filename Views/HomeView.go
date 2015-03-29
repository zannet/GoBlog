package Views

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HomeView(c *gin.Context) {
	obj := c.MustGet("data").(gin.H)
	c.HTML(http.StatusOK, "home", obj)
}