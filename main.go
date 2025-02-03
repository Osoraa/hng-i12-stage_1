package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type nomba struct {
	Number     string   `json:"number"`
	Prime      bool     `json:"is_prime"`
	Perfect    bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	Sum        int      `json:"price"`
	Fact       string   `json:"fun_fact"`
}
