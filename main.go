package main

import (
	"github.com/gin-gonic/gin"
)

type CodeFormat struct {
	Language string `json:"language" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Input    string `json:"input" binding:"required"`
}

func CodeHandler(c *gin.Context) {
	var userInput CodeFormat

	err := c.ShouldBindJSON(&userInput)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"language": userInput.Language,
		"code":     userInput.Code,
		"input":    userInput.Input,
	})

}

func main() {
	//start the route
	r := gin.Default()

	//send the request
	r.POST("/run", CodeHandler)

	r.Run()
}
