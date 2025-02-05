package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lcv-back/go-crm-api-shopdev/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2025")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.PUT("/user/1", controller.NewUserController().GetUserByID)
	}

	return r
}
