package config

import (
	"user-service/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}
func NewValidator() *validator.Validate {
	validator := validator.New()
	return validator
}
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandling())
	r.Use(middleware.MetricHttpRequest())
	r.Use(middleware.MetricHttpLatency())
	return r
}
