package practise

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// User 类型复用 GinDBUser.go 里的定义

func createUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User

		// 接收客户端发送的 JSON
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "请求数据格式错误",
			})
			return
		}

		// 将用户信息插入 MySQL
		result, err := db.Exec(
			"INSERT INTO users (name, age) VALUES (?, ?)",
			user.Name,
			user.Age,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "新增用户失败",
			})
			return
		}

		// 获取数据库生成的自增 ID
		id, err := result.LastInsertId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "获取用户 ID 失败",
			})
			return
		}

		user.ID = id

		// 返回新增后的用户信息
		c.JSON(http.StatusCreated, user)
	}
}

func GinDBInsertDemo() {
	db, err := sql.Open(
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/test",
	)
	if err != nil {
		fmt.Println("打开数据库失败:", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}

	r := gin.Default()

	r.POST("/user", createUserHandler(db))

	r.Run(":8095")
}
