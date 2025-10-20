package main

import (
	docs "user-service/docs"
	"user-service/src/config"
	"user-service/src/config/env"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User Service API
// @version 1.0
// @description This is a user service API server
// @termsOfService http://swagger.io/terms/
// @BasePath /api/users
// @schemes http https
func main() {
	env.NewEnv()
	logger := config.NewLogger()
	db := config.NewSql()
	validation := config.NewValidator()
	router := config.NewRouter()
	config.Setup(&config.DependenciesConfig{
		DB:         db,
		Logger:     logger,
		Validation: validation,
		Router:     router,
	})

	docs.SwaggerInfo.BasePath = "/api/users"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := router.Run(":4000")
	if err != nil {
		logger.Fatalf("failed run server: %s", err.Error())
	}
}
