package router

import (
	"gin-socmed/config"
	"gin-socmed/handler"
	"gin-socmed/middleware"
	"gin-socmed/repository"
	"gin-socmed/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/tweets")

	r.Use(middleware.JWTMiddleware())

	r.POST("/", postHandler.Create)
}
