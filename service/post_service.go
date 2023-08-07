package service

import (
	"gin-socmed/dto"
	"gin-socmed/entity"
	"gin-socmed/errorhandler"
	"gin-socmed/repository"
	"math"
	"time"

	"github.com/go-playground/validator/v10"
)

type PostService interface {
	FindAll(params *dto.FilterParams) (*[]dto.PostsResponse, *dto.Paginate, error)
	Detail(id int) (*dto.PostsResponse, error)
	Create(req *dto.PostRequest) error
	Update(id int, req *dto.PostRequest) error
	Delete(id int, userID int) error
	MyTweet(userID int, params *dto.FilterParams) (*[]dto.PostsResponse, *dto.Paginate, error)
}

type postService struct {
	repository repository.PostRepository
	validator  *validator.Validate
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
		validator:  validator.New(),
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
	if err := s.validator.Struct(req); err != nil {
		return &errorhandler.BadRequestError{Message: err.Error()}
	}

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

func (s *postService) Update(id int, req *dto.PostRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return &errorhandler.BadRequestError{Message: err.Error()}
	}

	currentPost, err := (s).Detail(id)
	if err != nil {
		return err
	}

	if req.UserID != currentPost.UserID {
		return &errorhandler.BadRequestError{Message: "it's not your tweet"}
	}

	post := entity.Post{
		Tweet:     req.Tweet,
		UpdatedAt: time.Now(),
	}

	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.Update(id, &post); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *postService) Delete(id int, userID int) error {
	currentPost, err := (s).Detail(id)
	if err != nil {
		return err
	}

	if userID != currentPost.UserID {
		return &errorhandler.BadRequestError{Message: "it's not your tweet"}
	}

	if err := s.repository.Delete(id); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *postService) MyTweet(userID int, params *dto.FilterParams) (*[]dto.PostsResponse, *dto.Paginate, error) {
	total, err := s.repository.TotalMyTweet(userID, params)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	posts, err := s.repository.AllMyTweet(userID, params)
	paginate := dto.Paginate{
		Total:      int(total),
		PerPage:    params.Limit,
		Page:       params.Page,
		TotalPages: int(math.Ceil(float64(total) / float64(params.Limit))),
	}

	return posts, &paginate, err
}
