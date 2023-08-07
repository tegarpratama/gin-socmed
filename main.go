package main

import (
	"fmt"
	"gin-socmed/config"
	"gin-socmed/dto"
	"gin-socmed/helper"
	"gin-socmed/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()

	// Configuration static file
	r.MaxMultipartMemory = 8 << 20
	r.Static("/public", "./public")

	api := r.Group("/api")

	api.GET("/check-health", func(c *gin.Context) {
		res := helper.Response(dto.ResponseParams{
			StatusCode: http.StatusOK,
			Message:    "It's work",
		})

		c.JSON(http.StatusOK, res)

		// Test error handler
		// errorhandler.HandleError(c, &errorhandler.NotFoundError{Message: "NotFoundError"})
		// errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: "InternalServerError"})
		// errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "BadRequestError"})
	})

	router.AuthRouter(api)
	router.PostRouter(api)
	router.UserRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}
