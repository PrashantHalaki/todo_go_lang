package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PrashantHalaki/todo_go_lang/models"
	"github.com/gin-gonic/gin"
)

type CreateTodoInput struct {
	Title string `json:"title" binding:"required"`
}

type UpdateTodoInput struct {
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	UpdatedAt time.Time `json:"UpdatedAt" gorm:"autoUpdateTime"`
}

type DeleteTodoInput struct {
	Status    int       `json:"status"`
	DeletedAt time.Time `json:"DeletedAt" gorm:"autoUpdateTime"`
}

// GET /todos
// Find all todos
func FindTodos(c *gin.Context) {
	var todos []models.Task
	models.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// GET /todos/:id
// Find a todo
func FindTodo(c *gin.Context) {
	// Get model if exist
	var todo models.Task
	models.DB.Where("id = ?", c.Param("id")).First(&todo)
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		fmt.Println("Errrrrrrr ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// POST /todo
// Create new task
func CreateTodo(c *gin.Context) {
	// Validate input
	var input CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task
	todo := models.Task{Title: input.Title, Completed: false}
	models.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// PATCH /todos/:id
// Update a task
func UpdateTodo(c *gin.Context) {
	// Get model if exist
	var todo models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Validate input
	var input UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&todo).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// DELETE /todos/:id
// Delete a task
func DeleteTodo(c *gin.Context) {
	// Get model if exist
	var todo models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&todo).Updates(DeleteTodoInput{Status: -1, DeletedAt: time.Now()})

	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted record!"})
}
