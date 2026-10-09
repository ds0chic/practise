package practise

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBExecDemo() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("打开数据库失败:", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(100), age INT)")
	if err != nil {
		fmt.Println("创建表失败:", err)
		return
	}

	result, err := db.Exec(
		"INSERT INTO users (name, age) VALUES (?, ?)",
		"Lucy",
		20,
	)
	if err != nil {
		fmt.Println("新增失败:", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("获取新增ID失败:", err)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("获取影响行数失败:", err)
		return
	}

	fmt.Println("新增用户ID:", id)
	fmt.Println("影响行数:", affected)
}
