# todo-api 📝

Go + Gin + GORM + SQLite による ToDo管理API ＋ Web GUI対応プロジェクトです。  
APIベースのタスク管理に加え、BootstrapベースのモダンなWeb UIを備えています。

---

## ✅ 機能概要

- [x] ToDoの登録／一覧表示／更新／削除（API + Web）
- [x] 完了フラグ（チェックボックス）対応
- [x] 編集モーダルによる更新（UX改善）
- [x] バリデーション（タイトル1〜30文字）
- [x] フラッシュメッセージ（更新・追加時）
- [x] BootstrapベースのレスポンシブUI

---

## 📁 ディレクトリ構成

todo-api/
├─ cmd/
│ └─ main.go # エントリポイント
├─ internal/
│ ├─ config/
│ │ └─ db.go # DB接続ロジック（SQLite）
│ ├─ handler/
│ │ ├─ api.go # REST APIハンドラ
│ │ └─ web.go # Webルーティング・操作ハンドラ
│ ├─ models/
│ │ └─ todo.go # Todoモデル（GORM）
│ └─ web/
│ ├─ static/
│ │ └─ style.css # Bootstrap拡張用スタイル
│ └─ templates/
│ └─ index.tmpl # Web UIテンプレート（HTML）
├─ go.mod / go.sum

---

## 🛠 使用技術

| 項目       | 使用技術                         |
|------------|----------------------------------|
| 言語       | Go 1.22+                          |
| Web FW     | Gin                              |
| DB ORM     | GORM（+ SQLite）                 |
| テンプレート | Go html/template                |
| UI         | Bootstrap 5 + モーダル／バリデーション |
| その他     | フラッシュメッセージ用 Cookie処理 |

---

## 🚀 起動方法
1. モジュール取得
go mod tidy

2. 起動
go run ./cmd/main.go

3. ブラウザでアクセス
http://localhost:8080/web