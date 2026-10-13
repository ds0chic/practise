package practise

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OrmUser 类型与 TableName 复用 OrmCreate.go 里的定义

func OrmOrderDemo() {
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

	// 如果表为空，添加测试数据
	var count int64
	if err := db.Model(&OrmUser{}).Count(&count).Error; err != nil {
		fmt.Println("统计失败:", err)
		return
	}

	if count == 0 {
		samples := []OrmUser{
			{Name: "Tom", Age: 18},
			{Name: "Jack", Age: 22},
			{Name: "Lucy", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Mike", Age: 21},
		}

		if err := db.Create(&samples).Error; err != nil {
			fmt.Println("插入失败:", err)
			return
		}
	}

	// 条件查询 + 降序排序
	var users []OrmUser

	result := db.Where("age > ?", 20).
		Order("age DESC").
		Find(&users)

	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
		return
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
