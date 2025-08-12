package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/shinichi.sunayama/todo-api/config"
	"github.com/shinichi.sunayama/todo-api/models"
	"github.com/shinichi.sunayama/todo-api/routes"
)

func main() {
	// .env ファイルから環境変数を読み込む
	// 例: DB_USER, DB_PASSWORD など
	err := godotenv.Load()
	if err != nil {
		// 読み込みに失敗した場合はログ出力してアプリケーションを停止
		log.Fatal("❌ .env ファイルの読み込みに失敗しました")
	}

	// データベースに接続（config パッケージ内で GORM を初期化）
	config.ConnectDatabase()

	// Todo モデルをもとにテーブルを自動作成・更新（マイグレーション）
	// 存在しない場合はテーブルを作成、定義変更も自動で反映
	config.DB.AutoMigrate(&models.Todo{})

	// Gin を使ったルーターを初期化
	r := routes.SetupRouter()

	// Web サーバーをポート 8080 で起動
	log.Println("✅ サーバー起動中：http://localhost:8080")
	r.Run(":8080")
}
