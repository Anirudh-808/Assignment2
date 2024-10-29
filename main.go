package main

import (
	"github.com/gin-gonic/gin"
)

type CodeFormat struct {
	language string `json:"language" binding:"required"`
	code     string `json:"code" binding:"required"`
	input    string `json:"input" binding:"required"`
}

func CodeHandler(c *gin.Context) {
	var userInput CodeFormat

	err := c.ShouldBindJSON(&userInput)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"py":    userInput.language,
		"code":  userInput.code,
		"input": userInput.input,
	})

}

func main() {
	//start the route
	r := gin.Default()

	//send the request
	r.POST("/run", CodeHandler)

	r.Run()
}
