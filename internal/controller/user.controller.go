package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lcv-back/go-crm-api-shopdev/internal/service"
	"github.com/lcv-back/go-crm-api-shopdev/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{"vile"})
}
