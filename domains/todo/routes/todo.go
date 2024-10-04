package routes

import (
	"github.com/99-66/go-gin-project-template/domains/todo/controllers"
	"github.com/99-66/go-gin-project-template/middlewares"
	"github.com/gin-gonic/gin"
)

func TodoRoute(r *gin.Engine, todo *controllers.TodoAPI) *gin.Engine {
	api := r.Group("/api")
	v1 := api.Group("/v1").Use(middlewares.TokenAuthMiddleware())
	v1.GET("/todo", todo.FindAll)
	v1.POST("/todo", todo.CreateOne)
	return r
}
