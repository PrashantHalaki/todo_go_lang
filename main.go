package main

import (
	"github.com/PrashantHalaki/todo_go_lang/controllers"
	"github.com/PrashantHalaki/todo_go_lang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/todos", controllers.FindTodos)
	r.GET("/todos/:id", controllers.FindTodo)
	r.POST("/todo", controllers.CreateTodo)
	r.PATCH("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	// Run the server
	r.Run()
}
