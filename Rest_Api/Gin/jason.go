package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.Default()

var message string
var nick string
	router.POST("/form_post", func(c *gin.Context) {
		message = c.PostForm("message")
		nick = c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.GET("/hhh", func(c *gin.Context) {
		fmt.Println(message)

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})//gorm
	})
	router.Run(":8080")
}
