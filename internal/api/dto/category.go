package dto

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
)

type CategoryResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   string    `json:"createdAt"`
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

func ToCreateCategoryModel(req CreateCategoryRequest) domain.Category {
	return domain.Category{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
	}
}

func ToUpdateCategoryModel(req UpdateCategoryRequest) domain.Category {
	return domain.Category{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
	}
}

func ToCategoryResponse(c domain.Category) CategoryResponse {
	return CategoryResponse{
		ID:          c.ID,
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
