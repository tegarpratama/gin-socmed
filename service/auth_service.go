package service

import (
	"gin-socmed/dto"
	"gin-socmed/entity"
	"gin-socmed/errorhandler"
	"gin-socmed/helper"
	"gin-socmed/repository"

	"github.com/go-playground/validator/v10"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
	validator  *validator.Validate
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
		validator:  validator.New(),
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	// custom validator
	genderValidator := func(fl validator.FieldLevel) bool {
		gender := fl.Field().String()
		return gender == "male" || gender == "female"
	}

	s.validator.RegisterValidation("gender", genderValidator)

	if err := s.validator.Struct(req); err != nil {
		return &errorhandler.BadRequestError{Message: err.Error()}
	}

	if emailExists := s.repository.EmailExists(req.Email); emailExists {
		return &errorhandler.BadRequestError{Message: "email already registered"}
	}

	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHash,
		Gender:   req.Gender,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	if err := s.validator.Struct(req); err != nil {
		return nil, &errorhandler.BadRequestError{Message: err.Error()}
	}

	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}

	return &data, nil
}
