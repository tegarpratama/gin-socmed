package service

import (
	"gin-socmed/dto"
	"gin-socmed/entity"
	"gin-socmed/errorhandler"
	"gin-socmed/repository"
	"math"
)

type PostService interface {
	FindAll(params *dto.FilterParams) (*[]dto.PostsResponse, *dto.Paginate, error)
	Detail(id int) (*dto.PostsResponse, error)
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

func (s *postService) FindAll(params *dto.FilterParams) (*[]dto.PostsResponse, *dto.Paginate, error) {
	total, err := s.repository.TotalData(params)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	posts, err := s.repository.FindAll(params)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	paginate := &dto.Paginate{
		Total:      int(total),
		PerPage:    params.Limit,
		Page:       params.Page,
		TotalPages: int(math.Ceil(float64(total) / float64(params.Limit))),
	}

	return posts, paginate, nil
}

func (s *postService) Detail(id int) (*dto.PostsResponse, error) {
	post, err := s.repository.Detail(&id)

	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "tweet not found"}
	}

	return &post, nil
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
