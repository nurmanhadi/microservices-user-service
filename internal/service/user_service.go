package service

import (
	"database/sql"
	"errors"
	"strings"
	"user-service/dto"
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(request dto.UserRequest) error
	LoginUser(request dto.UserRequest) (*dto.LoginResponse, error)
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
func (s *userService) RegisterUser(request dto.UserRequest) error {
	if err := s.validation.Struct(&request); err != nil {
		s.logger.WithError(err).Warn("validation failed during user registration")
		return err
	}
	newEmail := strings.ToLower(request.Email)
	totalUser, err := s.userRepository.CountByEmail(newEmail)
	if err != nil {
		s.logger.WithError(err).Error("failed to count user by email")
		return err
	}
	if totalUser > 0 {
		s.logger.WithField("email", newEmail).Warn("email already exists")
		return response.Except(400, "email already exists")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.WithError(err).Error("failed to generate hash password")
		return err
	}
	newID := uuid.NewString()
	user := &model.User{
		Id:       newID,
		Email:    newEmail,
		Password: string(hashPassword),
	}
	if err := s.userRepository.Insert(*user); err != nil {
		s.logger.WithError(err).Error("failed to insert new user")
		return err
	}
	s.logger.WithFields(logrus.Fields{
		"user_id": newID,
		"email":   newEmail,
	}).Info("user registered successfully")
	return nil
}
func (s *userService) LoginUser(request dto.UserRequest) (*dto.LoginResponse, error) {
	if err := s.validation.Struct(&request); err != nil {
		s.logger.WithError(err).Warn("validation failed during user login")
		return nil, err
	}
	newEmail := strings.ToLower(request.Email)
	user, err := s.userRepository.FindByEmail(newEmail)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.WithError(err).Warn("email or password wrong")
			return nil, response.Except(400, "email or password wrong")
		}
		s.logger.WithError(err).Error("failed to user find by email")
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		s.logger.WithError(err).Warn("email or password wrong")
		return nil, response.Except(400, "email or password wrong")
	}
	resp := &dto.LoginResponse{
		UserID: user.Id,
	}
	s.logger.WithField("user_id", user.Id).Info("user successfuly login")
	return resp, nil
}
