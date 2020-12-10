package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", homePage)
	router.GET("/api/v1/member", getReadMember)

	router.Run(":5000")
}

func homePage(c *gin.Context) {
	c.AsciiJSON(http.StatusOK, map[string]interface{}{
		"Message": "Success",
	})
}
