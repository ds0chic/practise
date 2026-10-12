package practise

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type OrmUser struct {
	ID   uint
	Name string
	Age  int
}

// TableName 让 AutoMigrate 建的仍是 users 表（与你原版 User 行为一致，避免建成 orm_users）。
func (OrmUser) TableName() string {
	return "users"
}

func OrmCreateDemo() {

	// 数据库连接信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"

	// 连接 MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}

	fmt.Println("数据库连接成功")

	// 根据 User 自动创建 users 表
	err = db.AutoMigrate(&OrmUser{})
	if err != nil {
		panic("创建表失败")
	}

	// 创建用户对象
	user := OrmUser{
		Name: "Tom",
		Age:  22,
	}

	// 插入数据库
	result := db.Create(&user)

	if result.Error != nil {
		panic("新增用户失败")
	}

	// 查看插入后的 ID
	fmt.Println("新增用户ID:", user.ID)

	// 查看影响行数
	fmt.Println("影响行数:", result.RowsAffected)
}
