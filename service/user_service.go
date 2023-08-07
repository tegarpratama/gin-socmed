package service

import (
	"gin-socmed/dto"
	"gin-socmed/errorhandler"
	"gin-socmed/repository"
	"math"
)

type UserService interface {
	GetUsers(params *dto.FilterParams) (*[]dto.UserResponse, *dto.Paginate, error)
	Detail(id int) (*dto.UserResponse, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) GetUsers(params *dto.FilterParams) (*[]dto.UserResponse, *dto.Paginate, error) {
	total, err := s.repository.TotalUsers(params)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	users, err := s.repository.GetUsers(params)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	paginate := &dto.Paginate{
		Total:      int(total),
		PerPage:    params.Limit,
		Page:       params.Page,
		TotalPages: int(math.Ceil(float64(total) / float64(params.Limit))),
	}

	return users, paginate, nil
}

func (s *userService) Detail(id int) (*dto.UserResponse, error) {
	post, err := s.repository.Detail(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: err.Error()}
	}

	return post, nil
}
