package practise

import (
	"github.com/gin-gonic/gin"
)

func GinUserDemo() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "tom",
			"age":  22,
		})
	})
	r.Run(":8087")
}
