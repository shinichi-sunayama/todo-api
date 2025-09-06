package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinichi.sunayama/todo-api/internal/models"
	"gorm.io/gorm"
)

func RenderIndex(c *gin.Context, db *gorm.DB) {
	var todos []models.Todo
	db.Order("created_at desc").Find(&todos)

	// フラッシュメッセージの取得とリセット
	flash, _ := c.Cookie("flash")
	c.SetCookie("flash", "", -1, "/", "localhost", false, false)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"todos": todos,
		"flash": flash,
	})
}

func ToggleDone(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var todo models.Todo

	if err := db.First(&todo, id).Error; err == nil {
		todo.Done = !todo.Done
		db.Save(&todo)
		c.SetCookie("flash", "完了状態を更新しました", 3, "/", "localhost", false, false)
	}

	c.Redirect(http.StatusSeeOther, "/web")
}

func DeleteFromWeb(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	db.Delete(&models.Todo{}, id)

	c.SetCookie("flash", "TODOを削除しました", 3, "/", "localhost", false, false)
	c.Redirect(http.StatusSeeOther, "/web")
}

func UpdateFromWeb(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	title := c.PostForm("title")

	if len(title) == 0 || len(title) > 30 {
		c.SetCookie("flash", "タイトルは1〜30文字で入力してください", 3, "/", "localhost", false, false)
		c.Redirect(http.StatusSeeOther, "/web")
		return
	}

	var todo models.Todo
	if err := db.First(&todo, id).Error; err == nil {
		todo.Title = title
		db.Save(&todo)
		c.SetCookie("flash", "TODOを更新しました", 3, "/", "localhost", false, false)
	}

	c.Redirect(http.StatusSeeOther, "/web")
}
