package config

import (
	"log"

	"github.com/shinichi.sunayama/todo-api/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// ✅ 自動マイグレーションの追加（←これがないとテーブルが作られない）
	err = database.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("マイグレーションに失敗しました: %v", err)
	}

	DB = database
	return database, nil
}
