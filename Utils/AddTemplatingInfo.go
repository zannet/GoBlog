package Utils

import (
	"github.com/gin-gonic/gin"
	"../Config"
)

func AddTemplatingInfo(c *gin.Context) {
	obj := c.MustGet("data").(gin.H)
	obj["_static_root"] = Config.Global.Service.Address + "/static"
	c.Set("data", obj)
}