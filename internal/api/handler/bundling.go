package handler

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/usecase"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/helper"
	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BundlingHandler struct {
	UseCase *usecase.BundlingUseCase
}

func NewBundlingHandler(u *usecase.BundlingUseCase) *BundlingHandler {
	return &BundlingHandler{UseCase: u}
}

func (h *BundlingHandler) Paginate(c *gin.Context) {
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

func (h *BundlingHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	bp, err := h.UseCase.FindById(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, bp, "success get bundling package")
}

func (h *BundlingHandler) Store(c *gin.Context) {
	var req dto.CreateBundlingPackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	created, err := h.UseCase.Store(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.Created(c, created, "success create bundling package")
}

func (h *BundlingHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	var req dto.UpdateBundlingPackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.Update(c.Request.Context(), id, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success update bundling package")
}

func (h *BundlingHandler) Activate(c *gin.Context) {
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

	response.OK(c, res, "message activate bundling package")
}

func (h *BundlingHandler) Deactivate(c *gin.Context) {
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

	response.OK(c, res, "message deactivate bundling package")
}
