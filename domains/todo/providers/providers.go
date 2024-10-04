package providers

import (
	"github.com/99-66/go-gin-project-template/domains/todo/controllers"
	"github.com/99-66/go-gin-project-template/domains/todo/repositories"
	"github.com/99-66/go-gin-project-template/domains/todo/services"
	"gorm.io/gorm"
)

func ProvideTodoAPI(t services.TodoService) *controllers.TodoAPI {
	return &controllers.TodoAPI{TodoService: t}
}

func ProvideTodoRepository(DB *gorm.DB) repositories.TodoRepository {
	return repositories.TodoRepository{DB: DB}
}

func ProvideTodoService(t repositories.TodoRepository) services.TodoService {
	return services.TodoService{TodoRepository: t}
}
