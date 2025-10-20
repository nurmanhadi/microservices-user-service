package routes

import (
	"user-service/src/rest/handler"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type RouteHandler struct {
	Router      *gin.Engine
	AuthHandler handler.AuthHandler
}

func (r *RouteHandler) Setup() {
	r.Router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	api := r.Router.Group("/api")

	user := api.Group("/users")
	auth := user.Group("/auth")
	auth.POST("/register", r.AuthHandler.RegisterUser)
	auth.POST("/login", r.AuthHandler.LoginUser)
}
