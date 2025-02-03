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
	router.GET("/api/classify-number/:num", getNumber)

	// Run server
	router.Use(cors.New(config))
	router.Run("localhost:8080")
}

// getNumbercalls the number_api
func getNumber(c *gin.Context) {
	

	num := c.DefaultQuery("num", "42")
	if num, err := strconv.Atoi(num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{{"number": "alphabet", "error": true}})

	// Return the response
	c.JSON(200, n)
}