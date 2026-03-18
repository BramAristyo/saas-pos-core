package handler

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/service"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ModifierGroupHandler struct {
	Service *service.ModifierGroupService
}

func NewModifierGroupHandler(s *service.ModifierGroupService) *ModifierGroupHandler {
	return &ModifierGroupHandler{
		Service: s,
	}
}

func (h *ModifierGroupHandler) Paginate(c *gin.Context) {
	var req filter.PaginationWithInputFilter
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.Service.Paginate(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OKPaginate(c, res.Data, res.Meta)
}

func (h *ModifierGroupHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.Service.FindById(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success get modifier group")
}

func (h *ModifierGroupHandler) Store(c *gin.Context) {
	var req dto.CreateModifierGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.Service.Store(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.Created(c, res, "success created modifier group")
}

func (h *ModifierGroupHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	var req dto.UpdateModifierGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := h.Service.Update(c.Request.Context(), id, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success update modifer group")
}

func (h *ModifierGroupHandler) Activate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.Service.UpdateStatus(c.Request.Context(), id, true)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success activate modifier group")
}

func (h *ModifierGroupHandler) Deactivate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.Service.UpdateStatus(c.Request.Context(), id, false)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success deactivate modifier group")
}
