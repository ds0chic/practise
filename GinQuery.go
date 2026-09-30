package practise

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginUserQueryHandler(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func GinQueryDemo() {
	r := gin.Default()

	r.GET("/user", ginUserQueryHandler)

	r.Run(":8088")
}
