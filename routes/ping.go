package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jessehorne/whousesgo.com/util"
	"net/http"
)

func GetPing(c *gin.Context) {
	version := util.GetVersion()

	if version == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "version not set on server",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"version": version,
	})
}
