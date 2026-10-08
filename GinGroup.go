package practise

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginLoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   "123456",
	})
}

func ginUserInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Tom",
		"age":  22,
	})
}

func ginUserOrderHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"order_id": "001",
	})
}

func GinGroupDemo() {
	r := gin.Default()

	// 第一层：所有接口统一以 /api 开头
	api := r.Group("/api")

	// 公开接口，不需要登录
	public := api.Group("/public")
	{
		public.POST("/login", ginLoginHandler)
	}

	// 用户接口
	user := api.Group("/user")

	// 只给 user 这一组添加登录中间件（复用 GinAuth.go 的 AuthMiddleware）
	user.Use(AuthMiddleware)

	{
		user.GET("/info", ginUserInfoHandler)
		user.GET("/order", ginUserOrderHandler)
	}

	r.Run(":8092")
}
