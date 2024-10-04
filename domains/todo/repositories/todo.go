package repositories

import (
	"github.com/99-66/go-gin-project-template/domains/todo/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func (r *TodoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	result := r.DB.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func (r *TodoRepository) CreateOne(todo models.Todo) (models.Todo, error) {
	result := r.DB.Create(&todo)
	if result.Error != nil {
		return models.Todo{}, result.Error
	}
	return todo, nil
}
