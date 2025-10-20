package main

import (
	"user-service/src/config"
	"user-service/src/config/env"
)

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

	err := router.Run(":4000")
	if err != nil {
		logger.Fatalf("failed run server: %s", err.Error())
	}
}
