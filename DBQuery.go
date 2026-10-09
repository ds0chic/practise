package practise

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBQueryDemo() {
	db, err := sql.Open(
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/test",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query(
		"SELECT id, name, age FROM users",
	)
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println("读取失败:", err)
			return
		}

		fmt.Println(id, name, age)
	}
}
