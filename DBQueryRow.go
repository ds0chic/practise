package practise

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBQueryRowDemo() {
	db, err := sql.Open(
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/test",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	var id int
	var name string
	var age int

	row := db.QueryRow(
		"SELECT id, name, age FROM users WHERE id = ?",
		1,
	)

	err = row.Scan(&id, &name, &age)
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}

	fmt.Println(id, name, age)
}
