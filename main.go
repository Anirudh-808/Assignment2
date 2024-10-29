/*
	NOTE:
	THE POST REQUEST IS RUN THROUGH POSTMAN

	IMPORTANT THING TO NOTE, THE "USER INPUT" IS TO BE GIVEN IN THE
	BODY OF THE INPUT SECTION IN POSTMAN(i.e THE SECTION ABOVE THE OUPUT PART)
*/

package main

import (
	"github.com/gin-gonic/gin"
)

// initializing the struct for userInput
type CodeFormat struct {
	/*
		NOTE:
		PREVIOUS ERROR AMMENDED
		Here remember to use capital letters for field names
		otherwise we cant "export" or use them anywhere else
	*/
	Language string
	Code     string
	Input    string
}

// the func for handling the request
// use c*gin.Context always in the handler functions
func CodeHandler(c *gin.Context) {
	//initializing the usesr input variable with CodeFormat type
	var userInput CodeFormat

	//the below binds the user input into the variable c using JSON format
	err := c.BindJSON(&userInput)

	//checks for any error
	if err != nil {
		//if any error ouput it
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//if no error give the required output
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
	r.POST("/code/run", CodeHandler)

	//run the route
	r.Run()
}
