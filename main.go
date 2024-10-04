package main

import (
	"log"

	"github.com/99-66/go-gin-project-template/config"
	"github.com/99-66/go-gin-project-template/consts"
	"github.com/99-66/go-gin-project-template/docs"
	"github.com/99-66/go-gin-project-template/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	var err error
	var dbHandler consts.DbHandlers

	dbHandler.DB, err = config.InitDB()
	if err != nil {
		panic(err)
	}
	sqlDB, _ := dbHandler.DB.DB()
	defer sqlDB.Close()

	r := initRoutes(dbHandler)
	log.Fatal(r.Run())
}

func initRoutes(dbHandler consts.DbHandlers) *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	// CORS allows all origins
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	r.Use(cors.New(conf))
	routes.RegisterRoutes(dbHandler, r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
