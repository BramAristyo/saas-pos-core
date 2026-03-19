package handler

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/usecase"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ModifierOptionHandler struct {
	UseCase *usecase.ModifierOptionUseCase
}

func NewModifierOptionHandler(u *usecase.ModifierOptionUseCase) *ModifierOptionHandler {
	return &ModifierOptionHandler{UseCase: u}
}

func (h *ModifierOptionHandler) Paginate(c *gin.Context) {
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

func (h *ModifierOptionHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.FindById(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success get modifier option")
}

func (h *ModifierOptionHandler) Store(c *gin.Context) {
	var req dto.CreateModifierOptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.Store(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.Created(c, res, "success created modifier option")
}

func (h *ModifierOptionHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	var req dto.UpdateModifierOptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.Update(c.Request.Context(), id, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success update modifier option")
}

func (h *ModifierOptionHandler) Activate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.UpdateStatus(c.Request.Context(), id, true)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success activate modifier option")
}

func (h *ModifierOptionHandler) Deactivate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.UseCase.UpdateStatus(c.Request.Context(), id, false)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success deactivate modifier option")
}
