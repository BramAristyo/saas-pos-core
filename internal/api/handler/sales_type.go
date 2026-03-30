package handler

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/usecase"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/helper"
	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/gin-gonic/gin"
)

type SalesTypeHandler struct {
	UseCase *usecase.SalesTypeUseCase
}

func NewSalesTypeHandler(u *usecase.SalesTypeUseCase) *SalesTypeHandler {
	return &SalesTypeHandler{UseCase: u}
}

func (h *SalesTypeHandler) Paginate(c *gin.Context) {
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

func (h *SalesTypeHandler) FindById(c *gin.Context) {
	id, err := helper.ParseUUID(c)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.FindById(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success get sales type")
}

func (h *SalesTypeHandler) Store(c *gin.Context) {
	var req dto.CreateSalesTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.Store(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.Created(c, res, "success create sales type")
}

func (h *SalesTypeHandler) Update(c *gin.Context) {
	id, err := helper.ParseUUID(c)
	if err != nil {
		c.Error(err)
		return
	}

	var req dto.UpdateSalesTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.Update(c.Request.Context(), id, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success update sales type")
}

func (h *SalesTypeHandler) Delete(c *gin.Context) {
	id, err := helper.ParseUUID(c)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.UseCase.Delete(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}

	response.OK(c, nil, "success delete sales type")
}

func (h *SalesTypeHandler) Restore(c *gin.Context) {
	id, err := helper.ParseUUID(c)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.Restore(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success restore sales type")
}
