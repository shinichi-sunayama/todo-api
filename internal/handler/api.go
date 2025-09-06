package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinichi.sunayama/todo-api/internal/models"
	"gorm.io/gorm"
)

func GetTodos(c *gin.Context, db *gorm.DB) {
	var todos []models.Todo
	db.Order("created_at desc").Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context, db *gorm.DB) {
	title := c.PostForm("title")
	if len(title) == 0 || len(title) > 30 {
		c.SetCookie("flash", "タイトルは1〜30文字で入力してください", 3, "/", "localhost", false, false)
		c.Redirect(http.StatusSeeOther, "/web")
		return
	}

	todo := models.Todo{Title: title}
	if err := db.Create(&todo).Error; err != nil {
		c.SetCookie("flash", "登録に失敗しました", 3, "/", "localhost", false, false)
		c.Redirect(http.StatusSeeOther, "/web")
		return
	}

	c.SetCookie("flash", "TODOを追加しました", 3, "/", "localhost", false, false)
	c.Redirect(http.StatusSeeOther, "/web")
}

func UpdateTodo(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var todo models.Todo
	if err := db.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "対象のTODOが見つかりません"})
		return
	}

	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不正な入力です"})
		return
	}

	db.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "削除に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "削除しました"})
}
