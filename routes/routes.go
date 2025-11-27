package routes

import (
	"github.com/dhanavadh/todol/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.GET("/todos", handlers.GetTodos)
	r.GET("/todos/:id", handlers.GetTodo)
	r.POST("/todos", handlers.CreateTodo)
	r.PUT("/todos/:id", handlers.UpdateTodo)
	r.DELETE("/todos/:id", handlers.DeleteTodo)
}
