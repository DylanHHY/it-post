package main

import (
	"it-post/database"
	"it-post/model"
)

func main() {
	database.ConnectToDB()

	// 执行 Post 模型的迁移
	err := database.DB.AutoMigrate(&model.Post{})
	if err != nil {
		panic("failed to migrate database")
	}
}
