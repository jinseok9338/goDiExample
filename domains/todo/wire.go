//go:build wireinject
// +build wireinject

package todo

import (
	"github.com/99-66/go-gin-project-template/domains/todo/controllers"
	"github.com/99-66/go-gin-project-template/domains/todo/providers"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func RegisterTodo(DB *gorm.DB) (*controllers.TodoAPI, error) {
	wire.Build(
		providers.ProvideTodoAPI,
		providers.ProvideTodoService,
		providers.ProvideTodoRepository,
	)
	return nil, nil
}
