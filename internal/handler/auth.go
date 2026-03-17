package handler

import (
	"github.com/BramAristyo/go-pos-mawish/internal/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/service"
	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: s,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.Service.Login(req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "login successfully")
}
