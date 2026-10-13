package practise

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OrmUser 类型与 TableName 复用 OrmCreate.go 里的定义

func OrmUpdateDemo() {
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

	user := OrmUser{Name: "Tom", Age: 22}

	if err := db.Create(&user).Error; err != nil {
		fmt.Println("新增失败:", err)
		return
	}

	fmt.Println("新增用户:", user)

	// 1. Update() 修改单个字段
	result := db.Model(&OrmUser{}).
		Where("id = ?", user.ID).
		Update("age", 23)

	if result.Error != nil {
		fmt.Println("修改失败:", result.Error)
		return
	}

	fmt.Println("影响行数:", result.RowsAffected)

	// 2. Updates() 使用 struct 修改多个字段
	result = db.Model(&OrmUser{}).
		Where("id = ?", user.ID).
		Updates(OrmUser{
			Name: "Tommy",
			Age:  0,
		})

	if result.Error != nil {
		fmt.Println("修改失败:", result.Error)
		return
	}

	var updated OrmUser
	if err := db.First(&updated, user.ID).Error; err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	fmt.Println("Struct 修改后:", updated)

	// 3. Updates() 使用 map 修改零值字段
	result = db.Model(&OrmUser{}).
		Where("id = ?", user.ID).
		Updates(map[string]any{
			"age": 0,
		})

	if result.Error != nil {
		fmt.Println("修改失败:", result.Error)
		return
	}

	var finalUser OrmUser
	if err := db.First(&finalUser, user.ID).Error; err != nil {
		fmt.Println("查询失败:", err)
		return
	}

	fmt.Println("最终用户:", finalUser)
}
