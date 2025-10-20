package dto

import (
	"time"
	"user-service/pkg/enum"
)

type UserRequest struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}
type LoginResponse struct {
	UserID string `json:"user_id"`
}
type UserName struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type UserResponse struct {
	Id        string      `json:"id"`
	Name      UserName    `json:"name"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	Role      enum.ROLE   `json:"role"`
	Status    enum.STATUS `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
type UserUpdateProfileRequest struct {
	FirstName *string `json:"first_name" validate:"omitempty,max=100"`
	LastName  *string `json:"last_name" validate:"omitempty,max=100"`
	Phone     *string `json:"phone" validate:"omitempty,max=20"`
}
type UserUpdateStatusRequest struct {
	Status enum.STATUS `json:"status" validate:"required,oneof=active inactive banned"`
}
