package handler

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/usecase"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/helper"
	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	UseCase *usecase.CategoryUseCase
}

func NewCategoryHandler(u *usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		UseCase: u,
	}
}

func (h *CategoryHandler) Paginate(c *gin.Context) {
	var req filter.PaginationWithInputFilter
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.Paginate(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OKPaginate(c, res.Data, res.Meta)
}

func (h *CategoryHandler) FindById(c *gin.Context) {
	id, err := helper.ParseUUID(c)
	if err != nil {
		c.Error(err)
		return
	}
	category, err := h.UseCase.FindById(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, category, "success get category")
}

func (h *CategoryHandler) Store(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	created, err := h.UseCase.Store(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.Created(c, created, "success create category")
}

func (h *CategoryHandler) Update(c *gin.Context) {
	var req dto.UpdateCategoryRequest
	id, err := helper.ParseUUID(c)
	if err != nil {
		c.Error(err)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.Update(c.Request.Context(), id, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success update category")
}

func (h *CategoryHandler) Activate(c *gin.Context) {
	id, err := helper.ParseUUID(c)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.UpdateStatus(c.Request.Context(), id, true)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success activate category")
}

func (h *CategoryHandler) Deactivate(c *gin.Context) {
	id, err := helper.ParseUUID(c)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.UpdateStatus(c.Request.Context(), id, false)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success deactivate category")
}
