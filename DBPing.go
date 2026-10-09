package practise

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBPingDemo() {
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

	fmt.Println("数据库连接成功")
}
