package main

import (
	"user-service/config"
	"user-service/database"
)

func main() {
	config.NewEnv()
	logger := config.NewLogger()
	db := database.NewSql()
	validation := config.NewValidator()
	router := config.NewRouter()
	config.Setup(&config.DependenciesConfig{
		DB:         db,
		Logger:     logger,
		Validation: validation,
		Router:     router,
	})

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		logger.Fatalf("failed run server: %s", err.Error())
	}
}
