package practise

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBUpdateDemo() {
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
		fmt.Println("数据库连接失败:", err)
		return
	}

	result, err := db.Exec(
		"UPDATE users SET age = ? WHERE id = ?",
		23,
		1,
	)
	if err != nil {
		fmt.Println("修改失败:", err)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("获取影响行数失败:", err)
		return
	}

	fmt.Println("影响行数:", affected)
}
