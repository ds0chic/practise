package practise

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func deleteUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "用户ID错误",
			})
			return
		}

		result, err := db.Exec(
			"DELETE FROM users WHERE id = ?",
			id,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "删除失败",
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

		if affected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "用户不存在",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "删除成功",
		})
	}
}

func GinDBDeleteDemo() {
	db, err := sql.Open(
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/test",
	)
	if err != nil {
		return
	}
	defer db.Close()

	r := gin.Default()

	r.DELETE("/user/:id", deleteUserHandler(db))

	r.Run(":8097")
}
