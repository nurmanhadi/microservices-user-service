package handler

import (
	"user-service/pkg/response"
	"user-service/src/dto"
	"user-service/src/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}
type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
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
func (h *userHandler) RegisterUser(ctx *gin.Context) {
	request := new(dto.UserRequest)
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.Error(response.Except(400, err.Error()))
		return
	}
	err := h.userService.RegisterUser(*request)
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
func (h *userHandler) LoginUser(ctx *gin.Context) {
	request := new(dto.UserRequest)
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.Error(response.Except(400, err.Error()))
		return
	}
	result, err := h.userService.LoginUser(*request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, result)
}
