package practise

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// User 类型复用 GinDBUser.go 里的定义

func updateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取要修改的用户 ID
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "用户 ID 不合法",
			})
			return
		}

		// 接收客户端发送的 JSON
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "请求数据格式错误",
			})
			return
		}

		// 修改数据库
		result, err := db.Exec(
			"UPDATE users SET name = ?, age = ? WHERE id = ?",
			user.Name,
			user.Age,
			id,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "修改失败",
			})
			return
		}

		affected, err := result.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "获取影响行数失败",
			})
			return
		}

		user.ID = int64(id)

		c.JSON(http.StatusOK, gin.H{
			"user":     user,
			"affected": affected,
		})
	}
}

func GinDBUpdateDemo() {
	db, err := sql.Open(
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/test",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	r := gin.Default()

	r.PUT("/user/:id", updateUserHandler(db))

	r.Run(":8096")
}
