package router

import (
	"gin-socmed/config"
	"gin-socmed/handler"
	"gin-socmed/repository"
	"gin-socmed/service"

	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	api.GET("/users", userHandler.Get)
	api.GET("/users/:id/detail", userHandler.Detail)
}
