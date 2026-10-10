package practise

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBDeleteDemo() {
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

	result, err := db.Exec(
		"DELETE FROM users WHERE id = ?",
		3,
	)
	if err != nil {
		fmt.Println("删除失败:", err)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("获取影响行数失败:", err)
		return
	}

	fmt.Println("删除行数:", affected)
}
