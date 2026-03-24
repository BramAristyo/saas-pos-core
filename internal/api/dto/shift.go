package dto

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ShiftExpenseResponse struct {
	ID          uuid.UUID                `json:"id"`
	ShiftID     uuid.UUID                `json:"shiftId"`
	Type        domain.ShiftExpensesType `json:"type"`
	Amount      decimal.Decimal          `json:"amount"`
	Description *string                  `json:"description"`
	CreatedAt   string                   `json:"createdAt"`
}

type ShiftResponse struct {
	ID            uuid.UUID              `json:"id"`
	OpenedBy      uuid.UUID              `json:"openedBy"`
	OpenedByUser  UserResponse           `json:"openedByUser"`
	ClosedBy      *uuid.UUID             `json:"closedBy,omitempty"`
	ClosedByUser  *UserResponse          `json:"closedByUser,omitempty"`
	OpeningCash   decimal.Decimal        `json:"openingCash"`
	ClosingCash   *decimal.Decimal       `json:"closingCash,omitempty"`
	Notes         *string                `json:"notes"`
	OpenedAt      string                 `json:"openedAt"`
	ClosedAt      *string                `json:"closedAt,omitempty"`
	CreatedAt     string                 `json:"createdAt"`
	ShiftExpenses []ShiftExpenseResponse `json:"shiftExpenses,omitempty"`
}

type ShiftResponsePagination struct {
	Data []ShiftResponse `json:"data"`
	Meta filter.Meta     `json:"meta"`
}

type OpenShiftRequest struct {
	OpeningCash decimal.Decimal `json:"openingCash" binding:"required"`
	Notes       *string         `json:"notes"`
}

type CloseShiftRequest struct {
	ClosingCash decimal.Decimal `json:"closingCash" binding:"required"`
	Notes       *string         `json:"notes"`
}

type ShiftExpenseRequest struct {
	ID          *string                  `json:"id"`
	Type        domain.ShiftExpensesType `json:"type" binding:"required,oneof=in out"`
	Amount      decimal.Decimal          `json:"amount" binding:"required"`
	Description *string                  `json:"description"`
}

type UpsertShiftExpensesRequest struct {
	Expenses []ShiftExpenseRequest `json:"expenses" binding:"required,dive"`
}

func ToShiftExpenseResponse(se *domain.ShiftExpenses) ShiftExpenseResponse {
	return ShiftExpenseResponse{
		ID:          se.ID,
		ShiftID:     se.ShiftID,
		Type:        se.Type,
		Amount:      se.Amount,
		Description: se.Description,
		CreatedAt:   se.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToShiftResponse(s *domain.Shift) ShiftResponse {
	var closedAt *string
	if s.ClosedAt != nil {
		str := s.ClosedAt.Format("2006-01-02 15:04:05")
		closedAt = &str
	}

	var closedByUser *UserResponse
	if s.ClosedByUser != nil {
		res := ToUserResponse(s.ClosedByUser)
		closedByUser = &res
	}

	expenses := make([]ShiftExpenseResponse, 0, len(s.ShiftExpenses))
	for _, e := range s.ShiftExpenses {
		expenses = append(expenses, ToShiftExpenseResponse(&e))
	}

	return ShiftResponse{
		ID:            s.ID,
		OpenedBy:      s.OpenedBy,
		OpenedByUser:  ToUserResponse(&s.OpenedByUser),
		ClosedBy:      s.ClosedBy,
		ClosedByUser:  closedByUser,
		OpeningCash:   s.OpeningCash,
		ClosingCash:   s.ClosingCash,
		Notes:         s.Notes,
		OpenedAt:      s.OpenedAt.Format("2006-01-02 15:04:05"),
		ClosedAt:      closedAt,
		CreatedAt:     s.CreatedAt.Format("2006-01-02 15:04:05"),
		ShiftExpenses: expenses,
	}
}

func ToShiftResponsePagination(s []ShiftResponse, p filter.PaginationWithInputFilter, totalRows int64) ShiftResponsePagination {
	return ShiftResponsePagination{
		Data: s,
		Meta: p.ToMeta(totalRows),
	}
}
