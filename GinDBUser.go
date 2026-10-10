package practise

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取路径参数
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "ID 必须是数字",
			})
			return
		}

		// 查询数据库
		var user User

		row := db.QueryRow(
			"SELECT id, name, age FROM users WHERE id = ?",
			id,
		)

		err = row.Scan(&user.ID, &user.Name, &user.Age)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "用户不存在",
			})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "数据库查询失败",
			})
			return
		}

		// 返回查询到的用户
		c.JSON(http.StatusOK, user)
	}
}

func GinDBUserDemo() {
	// 连接数据库
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

	// 创建 Gin 服务
	r := gin.Default()

	// 根据 ID 查询用户
	r.GET("/user/:id", getUserHandler(db))

	r.Run(":8094")
}
