package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinichi.sunayama/todo-api/config"
	"github.com/shinichi.sunayama/todo-api/models"
)

// GET /todos
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	if err := config.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "一覧取得に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// GET /todos/:id（必要なら）
func GetTodoByID(c *gin.Context) {
	var todo models.Todo
	if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todoが見つかりません"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// POST /todos
func CreateTodo(c *gin.Context) {
	var in models.Todo
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&in).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "作成に失敗しました"})
		return
	}
	c.JSON(http.StatusCreated, in)
}

// PUT /todos/:id
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todoが見つかりません"})
		return
	}
	var in models.Todo
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.Title = in.Title
	todo.Completed = in.Completed
	if err := config.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// DELETE /todos/:id（必要なら）
func DeleteTodo(c *gin.Context) {
	if err := config.DB.Delete(&models.Todo{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "削除に失敗しました"})
		return
	}
	c.Status(http.StatusNoContent)
}
