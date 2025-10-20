package dto

type UserRequest struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}
type LoginResponse struct {
	UserID string `json:"user_id"`
}
