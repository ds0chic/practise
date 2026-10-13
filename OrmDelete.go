package practise

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OrmUser 类型与 TableName 复用 OrmCreate.go 里的定义

func OrmDeleteDemo() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	if err := db.AutoMigrate(&OrmUser{}); err != nil {
		fmt.Println("创建表失败:", err)
		return
	}

	// 新增测试用户
	user := OrmUser{Name: "Tom", Age: 22}

	if err := db.Create(&user).Error; err != nil {
		fmt.Println("新增失败:", err)
		return
	}

	fmt.Println("新增用户:", user)

	// 根据 ID 删除用户
	result := db.Delete(&OrmUser{}, user.ID)

	if result.Error != nil {
		fmt.Println("删除失败:", result.Error)
		return
	}

	fmt.Println("删除行数:", result.RowsAffected)
}
