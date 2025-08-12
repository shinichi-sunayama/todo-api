package models

// Todo は「やることリスト（ToDo）」の 1 件を表すデータ構造です。
// GORM の ORM 機能を活用して、MySQL の todos テーブルとマッピングされます。
type Todo struct {
	ID        uint   `gorm:"primaryKey" json:"id"` // 主キー（自動インクリメント）
	Title     string `json:"title"`                // タイトル（ToDo の内容）
	Completed bool   `json:"completed"`            // 完了済みかどうか（true = 完了）
}
