package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	router.GET("/api/v1/member", getMember)
	router.GET("/api/v1/member/:id", getIdMember)
	router.POST("/api/v1/member", getCreateMember)
	router.PUT("/api/v1/member/:id", getEditMember)
	router.DELETE("/api/v1/member/:id", getDeleteMember)

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
func getIdMember(c *gin.Context) {
	id := c.Param("id")

	for i, u := range member {
		if u.Id == id {
			c.JSON(200, member[i])
		}
	}
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

	reqBody.Id = uuid.New().String() //Menggenerate ID secara otomatis

	member = append(member, reqBody)

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

func getEditMember(c *gin.Context) {
	id := c.Param("id")
	var reqBody Member
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}
	for index, idMember := range member {
		if idMember.Id == id {
			member[index].CodeMember = reqBody.CodeMember
			member[index].NickName = reqBody.NickName
			member[index].LastName = reqBody.LastName
			member[index].Status = reqBody.Status

			c.JSON(200, map[string]interface{}{
				"message": "Success",
			})
			return
		}
		c.JSON(404, gin.H{
			"error":   true,
			"message": "invalid to edit",
		})
	}
}

func getDeleteMember(c *gin.Context) {
	id := c.Param("id")

	for i, u := range member {
		if u.Id == id {
			member = append(member[:i], member[i+1:]...)

			c.JSON(200, map[string]interface{}{
				"message": "success",
			})
			return

		}
	}
	c.JSON(404, gin.H{
		"error":   true,
		"Message": "invalid delete member data",
	})
}
