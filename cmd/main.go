package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shinichi-sunayama/todo-api/internal/config"
	"github.com/shinichi-sunayama/todo-api/internal/handler"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("データベース接続に失敗しました: %v", err)
	}

	r := gin.Default()
	r.Static("/static", "./internal/web/static")
	r.LoadHTMLGlob("internal/web/templates/*.tmpl")

	api := r.Group("/api")
	{
		api.GET("/todos", func(c *gin.Context) {
			handler.GetTodos(c, db)
		})
		api.POST("/todos", func(c *gin.Context) {
			handler.CreateTodo(c, db)
		})
		api.PUT("/todos/:id", func(c *gin.Context) {
			handler.UpdateTodo(c, db)
		})
		api.DELETE("/todos/:id", func(c *gin.Context) {
			handler.DeleteTodo(c, db)
		})
	}

	r.GET("/web", func(c *gin.Context) {
		handler.RenderIndex(c, db)
	})
	r.POST("/toggle/:id", func(c *gin.Context) {
		handler.ToggleDone(c, db)
	})
	r.POST("/delete/:id", func(c *gin.Context) {
		handler.DeleteFromWeb(c, db)
	})
	r.POST("/update/:id", func(c *gin.Context) {
		handler.UpdateFromWeb(c, db)
	})

	r.Run(":8080")
}
