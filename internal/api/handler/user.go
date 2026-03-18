package handler

import (
	"fmt"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/service"
	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.Service.GetAll()
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, users, "success get users")
}

func (h *UserHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Service.FindById(id)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, user, "success get user")
}

func (h *UserHandler) Store(c *gin.Context) {
	var user dto.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(err)
		return
	}

	created, err := h.Service.Store(user)
	if err != nil {
		c.Error(err)
		return
	}

	response.Created(c, created, "success create user")
}

func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	fmt.Println(id, req.Email)
	updated, err := h.Service.Update(id, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, updated, "success update user")
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Destroy(id); err != nil {
		c.Error(err)
		return
	}

	response.OK(c, nil, "success delete user")
}
