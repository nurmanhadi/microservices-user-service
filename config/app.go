package config

import (
	"user-service/handler"
	"user-service/internal/repository"
	"user-service/internal/service"
	"user-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DependenciesConfig struct {
	DB         *sqlx.DB
	Logger     *logrus.Logger
	Validation *validator.Validate
	Router     *gin.Engine
}

func Setup(deps *DependenciesConfig) {
	// repository
	userRepo := repository.NewUserRepository(deps.DB)

	// service
	userServ := service.NewUserService(deps.Logger, deps.Validation, userRepo)

	// handler
	userHand := handler.NewUserHandler(userServ)

	// routes
	route := &routes.RouteHandler{
		Router:      deps.Router,
		UserHandler: userHand,
	}
	route.Setup()
}
