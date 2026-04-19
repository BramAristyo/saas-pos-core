package dto

import (
	"time"

	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ExpenseResponse struct {
	ID          uuid.UUID              `json:"id"`
	Category    domain.ExpenseCategory `json:"category"`
	Amount      decimal.Decimal        `json:"amount"`
	Description string                 `json:"description"`
	Date        string                 `json:"date"`
	DeletedAt   *string                `json:"deletedAt,omitempty"`
	CreatedAt   string                 `json:"createdAt"`
}

type ExpenseResponsePagination struct {
	Data []ExpenseResponse `json:"data"`
	Meta filter.Meta       `json:"meta"`
}

type CreateExpenseRequest struct {
	Category    domain.ExpenseCategory `json:"category" binding:"required"`
	Amount      decimal.Decimal        `json:"amount" binding:"required,gt=0"`
	Description string                 `json:"description" binding:"required"`
	Date        string                 `json:"date" binding:"required" example:"2006-01-02"`
}

type UpdateExpenseRequest struct {
	Category    domain.ExpenseCategory `json:"category" binding:"required"`
	Amount      decimal.Decimal        `json:"amount" binding:"required,gt=0"`
	Description string                 `json:"description" binding:"required"`
	Date        string                 `json:"date" binding:"required" example:"2006-01-02"`
}

func ToExpenseModel(req *CreateExpenseRequest) (domain.Expense, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return domain.Expense{}, err
	}
	return domain.Expense{
		Category:    req.Category,
		Amount:      req.Amount,
		Description: req.Description,
		Date:        date,
	}, nil
}

func ToUpdateExpenseModel(req *UpdateExpenseRequest) (domain.Expense, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return domain.Expense{}, err
	}
	return domain.Expense{
		Category:    req.Category,
		Amount:      req.Amount,
		Description: req.Description,
		Date:        date,
	}, nil
}

func ToExpenseResponse(e *domain.Expense) ExpenseResponse {
	resp := ExpenseResponse{
		ID:          e.ID,
		Category:    e.Category,
		Amount:      e.Amount,
		Description: e.Description,
		Date:        e.Date.Format("2006-01-02"),
		CreatedAt:   e.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if e.DeletedAt.Valid {
		t := e.DeletedAt.Time.Format("2006-01-02 15:04:05")
		resp.DeletedAt = &t
	}

	return resp
}

func ToExpenseResponsePagination(e []ExpenseResponse, p filter.PaginationWithInputFilter, totalRows int64) ExpenseResponsePagination {
	return ExpenseResponsePagination{
		Data: e,
		Meta: p.ToMeta(totalRows),
	}
}
