package practise

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginUserParamHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"user_id": id,
	})
}

func GinParamDemo() {
	r := gin.Default()

	r.GET("/user/:id", ginUserParamHandler)

	r.Run(":8089")
}
