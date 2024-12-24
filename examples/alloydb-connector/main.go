package main

import (
	"log"

	alloydbconnector "github.com/LinJianan2023/Polyhedra-Simple-Library-Golang/alloydb-connector"
)

func main() {
	// Initialize database connection
	// 初始化数据库连接
	if err := alloydbconnector.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Get GORM instance
	// 获取 GORM 实例
	db := alloydbconnector.GetDB()

	// Example: Define a user model
	// 示例：定义用户模型
	type User struct {
		ID   uint   `gorm:"primarykey"`
		Name string `gorm:"type:varchar(100)"`
	}

	// Auto migrate table structure
	// 自动迁移表结构
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	// Create user
	// 创建用户
	user := User{Name: "test user"}
	if err := db.Create(&user).Error; err != nil {
		log.Printf("Failed to create user: %v", err)
	}

	// Query users
	// 查询用户
	var users []User
	if err := db.Find(&users).Error; err != nil {
		log.Printf("Failed to query users: %v", err)
	}

	log.Printf("Found %d users", len(users))
}
