package dto

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TaxResponse struct {
	ID         uuid.UUID       `json:"id"`
	Name       string          `json:"name"`
	Percentage decimal.Decimal `json:"percentage"`
	IsActive   bool            `json:"isActive"`
	CreatedAt  string          `json:"createdAt"`
}

type TaxResponsePagination struct {
	Data []TaxResponse `json:"data"`
	Meta filter.Meta   `json:"meta"`
}

type CreateTaxRequest struct {
	Name       string          `json:"name" binding:"required,min=3,max=100"`
	Percentage decimal.Decimal `json:"percentage" binding:"required"`
}

type UpdateTaxRequest struct {
	Name       string          `json:"name" binding:"required,min=3,max=100"`
	Percentage decimal.Decimal `json:"percentage" binding:"required"`
	IsActive   bool            `json:"isActive"`
}

func ToCreateTaxModel(req *CreateTaxRequest) domain.Tax {
	return domain.Tax{
		Name:       req.Name,
		Percentage: req.Percentage,
		IsActive:   true,
	}
}

func ToUpdateTaxModel(req *UpdateTaxRequest) domain.Tax {
	return domain.Tax{
		Name:       req.Name,
		Percentage: req.Percentage,
		IsActive:   true,
	}
}

func ToTaxResponse(t *domain.Tax) TaxResponse {
	return TaxResponse{
		ID:         t.ID,
		Name:       t.Name,
		Percentage: t.Percentage,
		IsActive:   t.IsActive,
		CreatedAt:  t.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToTaxResponses(ts []domain.Tax) []TaxResponse {
	tsRes := make([]TaxResponse, len(ts))
	for i, t := range ts {
		tsRes[i] = ToTaxResponse(&t)
	}

	return tsRes
}

func ToTaxResponsePagination(t []TaxResponse, p filter.PaginationWithInputFilter, totalRows int64) TaxResponsePagination {
	return TaxResponsePagination{
		Data: t,
		Meta: p.ToMeta(totalRows),
	}
}
