package config

import (
	"fmt"
	"log"
	"time"
	"user-service/src/config/env"
	"user-service/src/rest/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func NewSql() *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.CONF.DB.Host,
		env.CONF.DB.Port,
		env.CONF.DB.Username,
		env.CONF.DB.Password,
		env.CONF.DB.Name,
	)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("failed connect to database: %s", err)
	}
	db.SetMaxIdleConns(env.CONF.DB.MaxIdleConns)
	db.SetMaxOpenConns(env.CONF.DB.MaxPoolConns)
	db.SetConnMaxLifetime(time.Duration(env.CONF.DB.MaxLifetime) * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("failed ping to database: %s", err)
	}
	return db
}
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
