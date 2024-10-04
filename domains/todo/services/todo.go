package services

import (
	"github.com/99-66/go-gin-project-template/domains/todo/models"
	"github.com/99-66/go-gin-project-template/domains/todo/repositories"
)

type TodoService struct {
	TodoRepository repositories.TodoRepository
}

func (s *TodoService) FindAll() ([]models.Todo, error) {
	return s.TodoRepository.FindAll()
}

func (s *TodoService) CreateOne(todo models.Todo) (models.Todo, error) {
	return s.TodoRepository.CreateOne(todo)
}
