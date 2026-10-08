package practise

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})

		c.Abort()
		return
	}

	c.Next()
}

func ginAuthUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Tom",
		"age":  22,
	})
}

func GinAuthDemo() {
	r := gin.Default()

	r.Use(AuthMiddleware)

	r.GET("/user", ginAuthUserHandler)

	r.Run(":8091")
}
