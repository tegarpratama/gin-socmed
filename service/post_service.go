package service

import (
	"gin-socmed/dto"
	"gin-socmed/entity"
	"gin-socmed/errorhandler"
	"gin-socmed/repository"
)

type PostService interface {
	Create(req *dto.PostRequest) error
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) Create(req *dto.PostRequest) error {
	post := entity.Post{
		UserID: req.UserID,
		Tweet:  req.Tweet,
	}

	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.Create(&post); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}
