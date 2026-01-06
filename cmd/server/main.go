package main

import (
	"net/http"

	"github.com/Ajaybalajiprasad/codeprofiles/pkg/fetcher"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/profile/:platform/:username", func(c *gin.Context) {
		platform := c.Param("platform")
		username := c.Param("username")

		data, err := fetcher.GetProfile(platform, username)
		
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    data,
		})
	})

	r.Run(":8080")
}