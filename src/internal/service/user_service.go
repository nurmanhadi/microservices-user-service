package service

import (
	"database/sql"
	"strings"
	"user-service/pkg/response"
	"user-service/src/dto"
	"user-service/src/internal/entity"
	"user-service/src/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	UpdateProfile(id string, request dto.UserUpdateProfileRequest) error
	GetUserByID(id string) (*dto.UserResponse, error)
	GetAllUsers() ([]dto.UserResponse, error)
	UpdateStatusByID(id string, request dto.UserUpdateStatusRequest) error
}
type userService struct {
	logger         *logrus.Logger
	validation     *validator.Validate
	userRepository repository.UserRepository
}

func NewUserService(logger *logrus.Logger, validation *validator.Validate, userRepository repository.UserRepository) UserService {
	return &userService{
		logger:         logger,
		validation:     validation,
		userRepository: userRepository,
	}
}
func (s *userService) UpdateProfile(id string, request dto.UserUpdateProfileRequest) error {
	if err := s.validation.Struct(&request); err != nil {
		s.logger.WithError(err).Warn("validation failed during update profile")
		return err
	}
	totalUser, err := s.userRepository.CountByID(id)
	if err != nil {
		s.logger.WithError(err).Error("failed to count by id")
		return err
	}
	if totalUser < 1 {
		s.logger.Warn("user not found")
		return response.Except(404, "user not found")
	}
	if request.Phone != nil {
		totalPhone, err := s.userRepository.CountByPhone(*request.Phone)
		if err != nil {
			s.logger.WithError(err).Error("failed to count by phone")
			return err
		}
		if totalPhone > 0 {
			s.logger.Warn("phone already exists")
			return response.Except(400, "phone already exists")
		}
	}
	user := &entity.User{
		FirstName: *request.FirstName,
		LastName:  *request.LastName,
		Phone:     *request.Phone,
	}
	if err := s.userRepository.UpdateProfile(id, *user); err != nil {
		s.logger.WithError(err).Error("failed to user update profile")
		return err
	}
	return nil
}
func (s *userService) GetUserByID(id string) (*dto.UserResponse, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.WithError(err).Warn("user not found")
			return nil, response.Except(404, "user not found")
		}
		s.logger.WithError(err).Error("failed to get user by id")
		return nil, err
	}

	resp := &dto.UserResponse{
		Id: user.Id,
		Name: dto.UserName{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Phone:     user.Phone,
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return resp, nil
}
func (s *userService) GetAllUsers() ([]dto.UserResponse, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		s.logger.WithError(err).Error("failed to get all users")
		return nil, err
	}

	responses := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, dto.UserResponse{
			Id: user.Id,
			Name: dto.UserName{
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
			Phone:     user.Phone,
			Email:     user.Email,
			Role:      user.Role,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return responses, nil
}
func (s *userService) UpdateStatusByID(id string, request dto.UserUpdateStatusRequest) error {
	if err := s.validation.Struct(&request); err != nil {
		s.logger.WithError(err).Warn("validation failed during update status")
		return err
	}
	totalUser, err := s.userRepository.CountByID(id)
	if err != nil {
		s.logger.WithError(err).Error("failed to count by id")
		return err
	}
	if totalUser < 1 {
		s.logger.Warn("user not found")
		return response.Except(404, "user not found")
	}
	newStatus := strings.ToLower(string(request.Status))
	if err := s.userRepository.UpdateStatusByID(id, newStatus); err != nil {
		s.logger.Errorf("failed to update user status: %v", err)
		return err
	}
	return nil
}
