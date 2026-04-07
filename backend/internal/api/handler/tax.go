package handler

import (
	"github.com/BramAristyo/saas-pos-core/backend/internal/api/dto"
	"github.com/BramAristyo/saas-pos-core/backend/internal/usecase"
	"github.com/BramAristyo/saas-pos-core/backend/pkg/filter"
	"github.com/BramAristyo/saas-pos-core/backend/pkg/helper"
	"github.com/BramAristyo/saas-pos-core/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type TaxHandler struct {
	UseCase *usecase.TaxUseCase
}

func NewTaxHandler(u *usecase.TaxUseCase) *TaxHandler {
	return &TaxHandler{
		UseCase: u,
	}
}

func (h *TaxHandler) Paginate(c *gin.Context) {
	var req filter.PaginationWithInputFilter
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(err)
		return
	}

	req.DynamicFilter.WithDefaultSort()

	res, err := h.UseCase.Paginate(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OKPaginate(c, res.Data, res.Meta)
}

func (h *TaxHandler) FindById(c *gin.Context) {
	id, err := helper.ParseUUID(c)
	if err != nil {
		c.Error(err)
		return
	}
	tax, err := h.UseCase.FindById(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, tax, "success get tax")
}

func (h *TaxHandler) Store(c *gin.Context) {
	var req dto.CreateTaxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	created, err := h.UseCase.Store(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.Created(c, created, "success create tax")
}

func (h *TaxHandler) Update(c *gin.Context) {
	var req dto.UpdateTaxRequest
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

	response.OK(c, res, "success update tax")
}

func (h *TaxHandler) Delete(c *gin.Context) {
	id, err := helper.ParseUUID(c)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.UseCase.Delete(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}

	response.OK(c, nil, "success delete tax")
}

func (h *TaxHandler) Restore(c *gin.Context) {
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

	response.OK(c, res, "success restore tax")
}
