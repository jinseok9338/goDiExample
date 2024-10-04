package controllers

import (
	"net/http"

	"github.com/99-66/go-gin-project-template/domains/todo/models"
	services "github.com/99-66/go-gin-project-template/domains/todo/services"
	"github.com/gin-gonic/gin"
)

type TodoAPI struct {
	TodoService services.TodoService
}

func (t *TodoAPI) FindAll(c *gin.Context) {
	todos, err := t.TodoService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (t *TodoAPI) CreateOne(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTodo, err := t.TodoService.CreateOne(newTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, createdTodo)
}
