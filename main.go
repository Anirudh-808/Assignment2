package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//start the route
	r := gin.Default()

	//send the request
	r.POST("/run", CodeHandler)

	r.Run()
}

type CodeFormat struct {
	language string
	code     string
	input    string
}

func CodeHandler(c *gin.Context) {
	var userInput CodeFormat

	c.BindJSON(&userInput)

	err := c.BindJSON(&userInput)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
	}

	c.JSON(200, gin.H{
		"py":    userInput.language,
		"code":  userInput.code,
		"input": userInput.input,
	})

}
