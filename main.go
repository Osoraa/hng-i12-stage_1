package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// number struct
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
	router.Run(":80")
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

	response, err := http.Get(fmt.Sprintf("http://numbersapi.com/%d/math", number))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"number": "alphabet", "error": true})
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"number": "alphabet", "error": true})
		return
	}

	fact.Number = number
	fact.Prime = isPrime(number)
	fact.Perfect = isPerfect(number)
	fact.Sum = getNumSum(number)
	fact.Fact = string(body)

	parity, armstrong := getProperties(number)
	fact.Properties = append(fact.Properties, parity)
	if armstrong {
		fact.Properties = append(fact.Properties, "armstrong")
	}

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

// Checks if a number is prime
func isPrime(num int) bool {
	//
	if num < 2 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true

}

// Checks if a number is perfect
func isPerfect(number int) bool {
	sum := 0

	for i := 1; i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}

	return sum == number
}

// Gets number properties
func getProperties(number int) (string, bool) {
	// Checks if a number is perfect
	var parity string
	armstrong := false

	// Get parity
	if number%2 == 0 {
		parity = "even"
	} else {
		parity = "odd"
	}

	// Check if armstrong
	sum := 0
	for i := number; i > 0; i /= 10 {
		sum += int(math.Pow(float64(i%10), 3))
	}

	if sum == number {
		armstrong = true
	}

	return parity, armstrong
}
