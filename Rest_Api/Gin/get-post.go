package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Username string `json:'username'`
	Password string `json:'password'`
}

var router *gin.Engine

func main() {
	router = gin.Default()
	initializeRoutes()
	router.Run(":8088")
}

func initializeRoutes() {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is index",
		})
		return
	})
	router.POST("/api", handleVerification)
	router.OPTIONS("/api", handleVerification)
	router.GET("/api", handleGet)
}

func handleGet(c *gin.Context) {
	message, e := c.GetQuery("m")
	fmt.Println(message)
	fmt.Println(e)

	c.String(http.StatusOK, "Get works! you sent: "+message)
}

func handleVerification(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		
		// setup headers
		c.Header("Allow", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
		c.Header("Content-Type", "application/json")
		c.Status(http.StatusOK)
	} else if c.Request.Method == "POST" {
		user := User{}
		err := c.BindJSON(&user)
		log.Println(err)
		log.Println(user)
		c.JSON(http.StatusOK, gin.H{
			"user": user.Username,
			"pass": user.Password,
		})
	}
}
