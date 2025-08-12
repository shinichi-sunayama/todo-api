package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shinichi.sunayama/todo-api/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Todoルートグループ
	todos := r.Group("/todos")
	{
		todos.GET("", controllers.GetTodos)       // 一覧取得
		todos.POST("", controllers.CreateTodo)    // 作成
		todos.PUT("/:id", controllers.UpdateTodo) // 更新（追加済み）
		// 必要に応じてDELETEなどもここに追加可能
	}

	return r
}
