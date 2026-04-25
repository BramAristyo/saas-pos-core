package handler

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/usecase"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"github.com/BramAristyo/saas-pos-core/server/pkg/response"
	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	usecase *usecase.AttendanceUseCase
}

func NewAttendanceHandler(u *usecase.AttendanceUseCase) *AttendanceHandler {
	return &AttendanceHandler{
		usecase: u,
	}
}

func (h *AttendanceHandler) Paginate(c *gin.Context) {
	var req filter.PaginationWithInputFilter
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.usecase.Paginate(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OKPaginate(c, res.Data, res.Meta)
}
