package handler

import (
	"fmt"
	"gin-socmed/dto"
	"gin-socmed/errorhandler"
	"gin-socmed/helper"
	"gin-socmed/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *postHandler {
	return &postHandler{
		service: service,
	}
}

func (h *postHandler) Get(c *gin.Context) {
	// 1. Without filter
	// page := c.DefaultQuery("page", "1")
	// limit := c.DefaultQuery("limit", "5")
	// search := c.Query("search")

	// pageNumber, _ := strconv.Atoi(page)
	// limitNumber, _ := strconv.Atoi(limit)
	// offset := (pageNumber - 1) * limitNumber

	// posts, paginate, err := h.service.FindAll(&dto.FilterParams{
	// 	Page:   pageNumber,
	// 	Limit:  limitNumber,
	// 	Offset: offset,
	// 	Search: search,
	// })

	// 2. With filter
	filter := helper.FilterParams(c)
	posts, paginate, err := h.service.FindAll(filter)

	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "List tweet's",
		Paginate:   paginate,
		Data:       posts,
	})

	c.JSON(http.StatusOK, res)
}

func (h *postHandler) Detail(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)
	post, err := h.service.Detail(idInt)

	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Detail tweet",
		Data:       post,
	})

	c.JSON(http.StatusOK, res)
}

func (h *postHandler) Create(c *gin.Context) {
	var post dto.PostRequest

	if err := c.ShouldBind(&post); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if post.Picture != nil {
		if err := os.MkdirAll("public/picture", 0755); err != nil {
			errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}

		// Rename picture
		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		// Save image into directory
		dst := filepath.Join("public/picture", filepath.Base(newFileName))
		c.SaveUploadedFile(post.Picture, dst)

		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", c.Request.Host, newFileName)
	}

	userID, _ := c.Get("userID")
	post.UserID = userID.(int)

	if err := h.service.Create(&post); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Succcess post your tweet",
	})

	c.JSON(http.StatusOK, res)
}

func (h *postHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	var post dto.PostRequest
	if err := c.ShouldBind(&post); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if post.Picture != nil {
		if err := os.MkdirAll("public/picture", 0755); err != nil {
			errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}

		// Rename picture
		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		// Save image into directory
		dst := filepath.Join("public/picture", filepath.Base(newFileName))
		c.SaveUploadedFile(post.Picture, dst)

		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", c.Request.Host, newFileName)
	}

	userID, _ := c.Get("userID")
	post.UserID = userID.(int)

	if err := h.service.Update(idInt, &post); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Successfully updated tweet",
	})

	c.JSON(http.StatusOK, res)
}

func (h *postHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	userID, _ := c.Get("userID")
	userIDInt := userID.(int)

	if err := h.service.Delete(idInt, userIDInt); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Tweet deleted",
	})

	c.JSON(http.StatusOK, res)
}

func (h *postHandler) MyTweet(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDInt := userID.(int)
	filter := helper.FilterParams(c)
	posts, paginate, err := h.service.MyTweet(userIDInt, filter)

	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "My Tweet",
		Paginate:   paginate,
		Data:       posts,
	})

	c.JSON(http.StatusOK, res)
}
