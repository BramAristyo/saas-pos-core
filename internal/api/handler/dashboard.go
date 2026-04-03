package handler

import (
	"github.com/BramAristyo/go-pos-mawish/internal/usecase"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	UseCase *usecase.DashboardUseCase
}

func NewDashboardHandler(u *usecase.DashboardUseCase) *DashboardHandler {
	return &DashboardHandler{UseCase: u}
}

func (h *DashboardHandler) SalesSummary(c *gin.Context) {
	var req filter.DynamicFilter
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(err)
		return
	}

	req.WithDefaultDateRange()

	res, err := h.UseCase.SalesSummary(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "Sales summary retrieved successfully")
}
