package handler

import (
	"user-service/pkg/response"
	"user-service/src/dto"
	"user-service/src/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}
type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

// RegisterUser godoc
// @Summary Register new user
// @Description Register a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "User registration data"
// @Success 201
// @Failure 400
// @Router /auth/register [post]
func (h *authHandler) RegisterUser(ctx *gin.Context) {
	request := new(dto.UserRequest)
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.Error(response.Except(400, err.Error()))
		return
	}
	err := h.authService.RegisterUser(*request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 201, "OK")
}

// LoginUser godoc
// @Summary Register user
// @Description Login a user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "User login data"
// @Success 200
// @Failure 400
// @Router /auth/login [post]
func (h *authHandler) LoginUser(ctx *gin.Context) {
	request := new(dto.UserRequest)
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.Error(response.Except(400, err.Error()))
		return
	}
	result, err := h.authService.LoginUser(*request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, result)
}
