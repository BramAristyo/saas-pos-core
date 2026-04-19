package handler

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/usecase"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"github.com/BramAristyo/saas-pos-core/server/pkg/helper"
	"github.com/BramAristyo/saas-pos-core/server/pkg/response"
	"github.com/gin-gonic/gin"
)

type ShiftExpensesHandler struct {
	UseCase *usecase.ShiftExpensesUseCase
}

func NewShiftExpensesHandler(u *usecase.ShiftExpensesUseCase) *ShiftExpensesHandler {
	return &ShiftExpensesHandler{
		UseCase: u,
	}
}

func (h *ShiftExpensesHandler) Paginate(c *gin.Context) {
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

func (h *ShiftExpensesHandler) FindById(c *gin.Context) {
	id, err := helper.ParseUUID(c)
	if err != nil {
		c.Error(err)
		return
	}
	expense, err := h.UseCase.FindById(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, expense, "success get shift expense")
}
