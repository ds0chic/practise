package practise

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OrmUser 类型与 TableName 复用 OrmCreate.go 里的定义

func OrmTxDemo() {
	// 1. 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&OrmUser{})

	// 2. 开启事务
	err = db.Transaction(func(tx *gorm.DB) error {
		// 插入第一个用户
		if err := tx.Create(&OrmUser{Name: "Tom", Age: 20}).Error; err != nil {
			return err
		}

		// 插入第二个用户
		if err := tx.Create(&OrmUser{Name: "Jack", Age: 22}).Error; err != nil {
			return err
		}

		return nil
	})

	// 3. 判断事务结果
	if err != nil {
		fmt.Println("事务失败，已回滚:", err)
	} else {
		fmt.Println("事务成功，已提交")
	}
}
