package routes

import (
	"user-service/handler"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type RouteHandler struct {
	Router      *gin.Engine
	UserHandler handler.UserHandler
}

func (r *RouteHandler) Setup() {
	r.Router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	api := r.Router.Group("/api")

	user := api.Group("/users")
	user.POST("/register", r.UserHandler.RegisterUser)
	user.POST("/login", r.UserHandler.LoginUser)
}
