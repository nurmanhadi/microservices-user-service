package handler

import (
	"user-service/dto"
	"user-service/internal/service"
	"user-service/pkg/response"

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
