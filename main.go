package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Member struct {
	Id         string `json:"id"`
	CodeMember string `json:"codeMember"`
	NickName   string `json:"nickName"`
	LastName   string `json:"lastName"`
	Status     bool   `json:"status"`
}

var member []Member

func main() {
	router := gin.Default()
	router.GET("/", homePage)
	router.GET("/api/v1/datamember", getMember)
	router.POST("/api/v1/member", getCreateMember)

	router.Run(":5000")
}

func homePage(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success",
	})
}

func getMember(c *gin.Context) {
	c.JSON(200, member)

}

func getCreateMember(c *gin.Context) {
	var reqBody Member
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}
	member = append(member, reqBody)

	c.JSON(200, gin.H{
		"message": "Success",
	})
}
