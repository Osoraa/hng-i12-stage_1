package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type nomba struct {
	Number     int      `json:"number"`
	Prime      bool     `json:"is_prime"`
	Perfect    bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	Sum        int      `json:"digit_sum"`
	Fact       string   `json:"fun_fact"`
}

// main function
func main() {
	// Create router
	router := gin.Default()

	// Handle CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	// GET request
	router.GET("/api/classify-number", getNumber)

	// Run server
	router.Use(cors.New(config))
	router.Run(":8081")
}

// getNumber assembles the nomba struct
func getNumber(c *gin.Context) {
	var fact nomba

	// num := c.Param("num")
	num := c.DefaultQuery("number", "42")
	number, err := strconv.Atoi(num)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"number": "alphabet", "error": true})
		return
	}
	
	fact.Number = number
	fact.Sum = getNumSum(number)

	// Return the response
	c.IndentedJSON(http.StatusOK, fact)
}

// Calculates sum of number's digits
func getNumSum(num int) int {
	sum := 0
	
	for num > 9 {
		sum += num % 10
		num /= 10
	}

	return sum + num
}

// 
// func getProperties(number int) []