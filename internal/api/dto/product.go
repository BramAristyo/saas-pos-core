package dto

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductResponse struct {
	ID          uuid.UUID         `json:"id"`
	CategoryID  uuid.UUID         `json:"category_id"`
	Category    *CategoryResponse `json:"category,omitempty"`
	Name        string            `json:"name"`
	Description *string           `json:"description"`
	Price       decimal.Decimal   `json:"price"`
	Cogs        decimal.Decimal   `json:"cogs"`
	ImageURL    *string           `json:"image_url"`
	IsActive    bool              `json:"is_active"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
}

type ProductResponsePagination struct {
	Data []ProductResponse `json:"data"`
	Meta filter.Meta       `json:"meta"`
}

type CreateProductRequest struct {
	CategoryID  uuid.UUID       `json:"category_id" binding:"required"`
	Name        string          `json:"name" binding:"required,min=3,max=100"`
	Description *string         `json:"description"`
	Price       decimal.Decimal `json:"price" binding:"required"`
	Cogs        decimal.Decimal `json:"cogs" binding:"required"`
	ImageURL    *string         `json:"image_url"`
}

type UpdateProductRequest struct {
	CategoryID  uuid.UUID       `json:"category_id" binding:"required"`
	Name        string          `json:"name" binding:"required,min=3,max=100"`
	Description *string         `json:"description"`
	Price       decimal.Decimal `json:"price" binding:"required"`
	Cogs        decimal.Decimal `json:"cogs" binding:"required"`
	ImageURL    *string         `json:"image_url"`
	IsActive    bool            `json:"is_active"`
}

type ChangeProductStatusRequest struct {
	IsActive bool `json:"is_active"`
}

func ToProductResponse(p domain.Product) ProductResponse {
	var category *CategoryResponse
	if p.Category != nil {
		c := ToCategoryResponse(*p.Category)
		category = &c
	}

	return ProductResponse{
		ID:          p.ID,
		CategoryID:  p.CategoryID,
		Category:    category,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Cogs:        p.Cogs,
		ImageURL:    p.ImageURL,
		IsActive:    p.IsActive,
		CreatedAt:   p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   p.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToProductResponsePagination(p []ProductResponse, f filter.PaginationWithInputFilter, totalRows int64) ProductResponsePagination {
	return ProductResponsePagination{
		Data: p,
		Meta: f.ToMeta(totalRows),
	}
}

func ToProductModel(req CreateProductRequest) domain.Product {
	return domain.Product{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Cogs:        req.Cogs,
		ImageURL:    req.ImageURL,
		IsActive:    true,
	}
}

func ToUpdateProductModel(req UpdateProductRequest) domain.Product {
	return domain.Product{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Cogs:        req.Cogs,
		ImageURL:    req.ImageURL,
		IsActive:    req.IsActive,
	}
}
