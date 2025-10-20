package config

import (
	"user-service/src/internal/repository"
	"user-service/src/internal/service"
	"user-service/src/rest/handler"
	"user-service/src/rest/routes"

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
	authServ := service.NewAuthService(deps.Logger, deps.Validation, userRepo)
	userServ := service.NewUserService(deps.Logger, deps.Validation, userRepo)

	// handler
	authHand := handler.NewAuthHandler(authServ)
	userHand := handler.NewUserHandler(userServ)

	// routes
	route := &routes.RouteHandler{
		Router:      deps.Router,
		AuthHandler: authHand,
		UserHandler: userHand,
	}
	route.Setup()
}
