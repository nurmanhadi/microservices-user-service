package routes

import (
	"user-service/pkg/enum"
	"user-service/src/rest/handler"
	"user-service/src/rest/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type RouteHandler struct {
	Router      *gin.Engine
	AuthHandler handler.AuthHandler
	UserHandler handler.UserHandler
}

func (r *RouteHandler) Setup() {
	r.Router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	api := r.Router.Group("/api")

	user := api.Group("/users")

	auth := user.Group("/auth")
	auth.POST("/register", r.AuthHandler.RegisterUser)
	auth.POST("/login", r.AuthHandler.LoginUser)

	profile := user.Group("/profiles")
	profile.GET("/",
		middleware.JwtValidation([]enum.ROLE{enum.ROLE_ADMIN}),
		r.UserHandler.GetAllUsers)
	profile.PUT("/:id",
		middleware.JwtValidation([]enum.ROLE{enum.ROLE_ADMIN, enum.ROLE_USER}),
		r.UserHandler.UpdateProfile)
	profile.GET("/:id",
		middleware.JwtValidation([]enum.ROLE{enum.ROLE_ADMIN, enum.ROLE_USER}),
		r.UserHandler.GetUserByID)
	profile.PUT("/:id/status",
		middleware.JwtValidation([]enum.ROLE{enum.ROLE_ADMIN}),
		r.UserHandler.UpdateStatusByID)
}
