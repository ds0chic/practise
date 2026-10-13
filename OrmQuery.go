package practise

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OrmUser 类型与 TableName 复用 OrmCreate.go 里的定义

func OrmQueryDemo() {
	// 1. 连接 MySQL
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}

	// 2. 自动创建表
	if err := db.AutoMigrate(&OrmUser{}); err != nil {
		fmt.Println("创建表失败:", err)
		return
	}

	// 3. 插入测试数据
	users := []OrmUser{
		{Name: "Tom", Age: 22},
		{Name: "Jack", Age: 25},
		{Name: "Lucy", Age: 20},
	}

	if err := db.Create(&users).Error; err != nil {
		fmt.Println("插入失败:", err)
		return
	}

	// 4. First() 查询一条
	var user OrmUser

	result := db.First(&user, users[0].ID)
	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
		return
	}

	fmt.Println("查询单条:", user)

	// 5. Find() 查询所有用户
	var allUsers []OrmUser

	result = db.Find(&allUsers)
	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
		return
	}

	fmt.Println("查询全部:", allUsers)

	// 6. Where() 条件查询
	var olderUsers []OrmUser

	result = db.Where("age > ?", 22).Find(&olderUsers)
	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
		return
	}

	fmt.Println("年龄大于22:", olderUsers)
}
