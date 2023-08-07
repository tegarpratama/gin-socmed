package handler

import (
	"gin-socmed/dto"
	"gin-socmed/errorhandler"
	"gin-socmed/helper"
	"gin-socmed/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *userHandler {
	return &userHandler{
		service: service,
	}
}

func (s *userHandler) Get(c *gin.Context) {
	filter := helper.FilterParams(c)
	users, paginate, err := s.service.GetUsers(filter)

	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "List user's",
		Paginate:   paginate,
		Data:       users,
	})

	c.JSON(http.StatusOK, res)
}

func (s *userHandler) Detail(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)
	user, err := s.service.Detail(idInt)

	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Detail user",
		Data:       user,
	})

	c.JSON(http.StatusOK, res)
}
