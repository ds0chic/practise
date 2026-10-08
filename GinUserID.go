package practise

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginUserIDHandler(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "userID not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"name":    "Tom",
	})
}

func GinUserIDDemo() {
	r := gin.Default()

	user := r.Group("/user")
	user.Use(AuthMiddleware)

	user.GET("/info", ginUserIDHandler)

	r.Run(":8093")
}
