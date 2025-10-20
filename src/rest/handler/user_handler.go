package handler

import (
	"user-service/pkg/response"
	"user-service/src/dto"
	"user-service/src/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UpdateProfile(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	UpdateStatusByID(ctx *gin.Context)
}
type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

// UpdateProfile godoc
// @Summary Update profile
// @Description Update user profile
// @Tags profiles
// @Accept json
// @Produce json
// @Param request body dto.UserUpdateProfileRequest true "User update profile data"
// @Param id path string true "User id"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /profiles/{id} [put]
func (h *userHandler) UpdateProfile(ctx *gin.Context) {
	id := ctx.Param("id")
	request := new(dto.UserUpdateProfileRequest)
	if err := ctx.ShouldBind(request); err != nil {
		ctx.Error(response.Except(400, "failed to parse json"))
		return
	}
	err := h.userService.UpdateProfile(id, *request)
	if err != nil {
		ctx.Error(err)
	}
	response.Success(ctx, 200, "OK")
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get a user detail by ID
// @Tags profiles
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 404
// @Router /profiles/{id} [get]
func (h *userHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.userService.GetUserByID(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, result)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all registered users
// @Tags profiles
// @Produce json
// @Success 200
// @Router /profiles [get]
func (h *userHandler) GetAllUsers(ctx *gin.Context) {
	result, err := h.userService.GetAllUsers()
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, result)
}

// UpdateStatusByID godoc
// @Summary Update user status by ID
// @Description Update user account status (active, inactive, banned)
// @Tags profiles
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body dto.UserUpdateStatusRequest true "User status data"
// @Success 200
// @Failure 400
// @Router /profiles/{id}/status [put]
func (h *userHandler) UpdateStatusByID(ctx *gin.Context) {
	id := ctx.Param("id")
	request := new(dto.UserUpdateStatusRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.Error(response.Except(400, "failed to parse json"))
		return
	}
	if err := h.userService.UpdateStatusByID(id, *request); err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, "OK")
}
