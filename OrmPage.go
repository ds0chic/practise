package practise

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OrmUser 类型与 TableName 复用 OrmCreate.go 里的定义

func OrmPageDemo() {
	// 1. 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	// 2. 自动创建表
	if err := db.AutoMigrate(&OrmUser{}); err != nil {
		fmt.Println("创建表失败:", err)
		return
	}

	// 3. 如果表为空，插入测试数据
	var existing int64
	if err := db.Model(&OrmUser{}).Count(&existing).Error; err != nil {
		fmt.Println("统计失败:", err)
		return
	}

	if existing == 0 {
		samples := []OrmUser{
			{Name: "Tom", Age: 18},
			{Name: "Jack", Age: 22},
			{Name: "Lucy", Age: 25},
			{Name: "Alice", Age: 19},
			{Name: "Bob", Age: 30},
			{Name: "Mike", Age: 21},
			{Name: "Rose", Age: 24},
			{Name: "John", Age: 23},
			{Name: "Anna", Age: 20},
			{Name: "David", Age: 26},
		}

		if err := db.Create(&samples).Error; err != nil {
			fmt.Println("插入失败:", err)
			return
		}
	}

	// 4. 设置查询和分页参数
	minAge := 20
	page := 2
	pageSize := 3
	offset := (page - 1) * pageSize

	// 5. 统计符合条件的总记录数
	var total int64

	err = db.Model(&OrmUser{}).
		Where("age > ?", minAge).
		Count(&total).Error

	if err != nil {
		fmt.Println("统计失败:", err)
		return
	}

	// 6. 条件查询 + 排序 + 分页
	var users []OrmUser

	result := db.Model(&OrmUser{}).
		Where("age > ?", minAge).
		Order("id DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&users)

	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
		return
	}

	// 7. 输出查询结果
	fmt.Println("符合条件的总人数:", total)
	fmt.Println("当前页:", page)
	fmt.Println("每页数量:", pageSize)

	for _, user := range users {
		fmt.Println(user)
	}
}
