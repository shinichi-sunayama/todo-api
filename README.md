# Todo API (Go × Gin × MySQL × Docker)

このプロジェクトは **Go + Gin + MySQL** を使用したシンプルな Todo API です。  
Docker Compose を使って、アプリケーションとデータベースを簡単に起動できます。

---

## 🚀 環境

- **言語**: Go 1.24
- **フレームワーク**: Gin
- **ORM**: Gorm
- **データベース**: MySQL 8.0
- **コンテナ管理**: Docker Compose

---

## 📂 ディレクトリ構成

├── config/ # DB接続設定
├── controllers/ # APIハンドラ
├── models/ # データベースモデル
├── routes/ # ルーティング定義
├── .env # 環境変数設定
├── docker-compose.yml
├── Dockerfile
├── go.mod / go.sum
└── main.go

## ⚙️ セットアップ手順

### 1. リポジトリのクローン
```bash
git clone https://github.com/yourname/todo-api.git
cd todo-api

2. .env の作成
以下を .env ファイルに記載（プロジェクトルート）

DB_USER=todo_user
DB_PASSWORD=todopass
DB_NAME=todo_db
DB_HOST=db
DB_PORT=3306

📡 APIエンドポイント
1. Todo一覧取得 (GET)
curl -X GET http://localhost:8080/todos

2. Todo作成 (POST)
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Gin", "completed":false}'

3. Todo更新 (PUT)
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated title", "completed":true}'

4. Todo削除 (DELETE)
curl -X DELETE http://localhost:8080/todos/1

🐳 Dockerコマンド例
コンテナ停止
docker compose down
再ビルド
docker compose up --build
ログ確認
docker logs todo-api-app
docker logs todo-api-db