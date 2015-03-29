package Middlewares

import (
	"github.com/gin-gonic/gin"
	"../Config"
)

func AddTemplatingInfo(c *gin.Context) {
	// Get the actual data from the context
	obj := c.MustGet("data").(gin.H)
	// Add the required data
	obj["_static_root"] = Config.Global.Service.Address + "/static"
	// Set the web's title
	obj["_site_title"] = "Sirikon's Lab"
	// And save it again
	c.Set("data", obj)
}