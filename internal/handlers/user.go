package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/BramAristyo/go-pos-mawish/internal/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/services"
	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/BramAristyo/go-pos-mawish/pkg/service_errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.Service.GetAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get user data")
		return
	}

	response.OK(c, users, "success get users")
}

func (h *UserHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Service.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Error(c, http.StatusNotFound, "failed to get user")
			return
		}
		response.Error(c, http.StatusInternalServerError, "failed to get user")
		return
	}

	response.OK(c, user, "success get user")
}

func (h *UserHandler) Store(c *gin.Context) {
	var user dto.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	created, err := h.Service.Store(user)
	if err != nil {
		var serviceErr *service_errors.ServiceError
		if errors.As(err, &serviceErr) {
			response.Error(c, serviceErr.Code, serviceErr.Message)
			return
		}
		response.Error(c, http.StatusInternalServerError, "failed to create user")
		return
	}

	response.Created(c, created, "success create user")
}

func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(id, req.Email)
	updated, err := h.Service.Update(id, req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Error(c, http.StatusNotFound, "failed to update user")
			return
		}

		response.Error(c, http.StatusInternalServerError, "failed to update user")
		return
	}

	response.OK(c, updated, "success update user")
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Destroy(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Error(c, http.StatusNotFound, "failed to delete user")
			return
		}
		response.Error(c, http.StatusInternalServerError, "failed to delete user")
		return
	}

	response.OK(c, nil, "success delete user")
}
