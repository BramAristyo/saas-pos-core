package dto

import (
	"github.com/BramAristyo/go-pos-mawish/internal/models"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
)

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
}

type CategoryResponsePagination struct {
	Data []CategoryResponse `json:"data"`
	Meta filter.Meta        `json:"meta"`
}

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required,min=6,max=100"`
	Description string `json:"description"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name" binding:"required,min=6,max=100"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}

func ToCreateCategoryModel(req CreateCategoryRequest) models.Category {
	return models.Category{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
	}
}

func ToUpdateCategoryModel(req UpdateCategoryRequest) models.Category {
	return models.Category{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
	}
}

func ToCategoryResponse(c models.Category) CategoryResponse {
	return CategoryResponse{
		ID:          int(c.ID),
		Name:        c.Name,
		Description: c.Description,
		IsActive:    c.IsActive,
		CreatedAt:   c.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToCategoryResponsePagination(c []CategoryResponse, p filter.PaginationWithInputFilter, totalRows int64) CategoryResponsePagination {
	return CategoryResponsePagination{
		Data: c,
		Meta: p.ToMeta(totalRows),
	}
}
