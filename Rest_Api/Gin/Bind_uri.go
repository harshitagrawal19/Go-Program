package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)
type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		fmt.Println(c.Param("id"))
		fmt.Println(c.Param("name"))
		if err := c.ShouldBindUri(&person); err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8088")
}
