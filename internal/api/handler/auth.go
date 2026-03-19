package handler

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/usecase"
	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UseCase *usecase.AuthUseCase
}

func NewAuthHandler(u *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		UseCase: u,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.Login(req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "login successfully")
}
