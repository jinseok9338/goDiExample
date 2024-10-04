package routes

import (
	"github.com/99-66/go-gin-project-template/consts"
	"github.com/99-66/go-gin-project-template/domains/todo"
	"github.com/99-66/go-gin-project-template/domains/todo/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(dbHandler consts.DbHandlers, server *gin.Engine) {
	RegisterTodoRoutes(dbHandler, server)
}

func RegisterTodoRoutes(dbHandler consts.DbHandlers, server *gin.Engine) {
	todoAPIController, err := todo.RegisterTodo(dbHandler.DB)
	if err != nil {
		panic(err)
	}
	routes.TodoRoute(server, todoAPIController)
}
