package practise

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginUserBindHandler(c *gin.Context) {
	type User struct {
		Name string `json:"name"`
		Age  string `json:"age"`
	}
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "格式错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
		"age":  user.Age,
	})
}

func GinBindDemo() {
	r := gin.Default()

	r.POST("/user", ginUserBindHandler)

	r.Run(":8090")
}
