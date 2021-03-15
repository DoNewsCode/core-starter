package handlers

import (
	"github.com/DoNewsCode/core-starter/app/repositories"
	"github.com/DoNewsCode/core/srvhttp"
	"github.com/gin-gonic/gin"
)

func NewUserHandler(repository repositories.UserRepository) *UserHandler {
	return &UserHandler{
		repository: repository,
	}
}

type UserHandler struct {
	repository repositories.UserRepository
}

func (h *UserHandler) GetOne(c *gin.Context) {
	encoder := srvhttp.NewResponseEncoder(c.Writer)
	encoder.EncodeResponse(h.repository.GetOne(1))
}
