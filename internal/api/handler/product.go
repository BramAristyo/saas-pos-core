package handler

import (
	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/service"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	Service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{
		Service: s,
	}
}

func (h *ProductHandler) Paginate(c *gin.Context) {
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

func (h *ProductHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}
	product, err := h.Service.FindById(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, product, "success get product")
}

func (h *ProductHandler) Store(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := h.bindJSON(c, &req); err != nil {
		return
	}

	created, err := h.Service.Store(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	response.Created(c, created, "success create product")
}

func (h *ProductHandler) Update(c *gin.Context) {
	var req dto.UpdateProductRequest
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}
	if err := h.bindJSON(c, &req); err != nil {
		return
	}

	res, err := h.Service.Update(c.Request.Context(), id, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success update product")
}

func (h *ProductHandler) ChangeStatus(c *gin.Context) {
	var req dto.ChangeProductStatusRequest
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(err)
		return
	}
	if err := h.bindJSON(c, &req); err != nil {
		return
	}

	res, err := h.Service.ChangeStatus(c.Request.Context(), id, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.OK(c, res, "success change product status")
}

func (h *ProductHandler) bindJSON(c *gin.Context, obj any) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.Error(err)
		return err
	}
	return nil
}
