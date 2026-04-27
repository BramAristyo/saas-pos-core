package dto

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CashTransactionResponse struct {
	ID          uuid.UUID       `json:"id"`
	CoaName     string          `json:"coaName"`
	Type        string          `json:"type"`
	Amount      decimal.Decimal `json:"amount"`
	Description *string         `json:"description"`
	Date        string          `json:"date"`
	CreatedAt   string          `json:"createdAt"`
}

type CashTransactionResponsePagination struct {
	Data []CashTransactionResponse `json:"data"`
	Meta filter.Meta               `json:"meta"`
}

func toCashTransactionResponse(ct *domain.CashTransaction) CashTransactionResponse {
	return CashTransactionResponse{
		ID:          ct.ID,
		CoaName:     ct.COA.Name,
		Type:        string(ct.Type),
		Amount:      ct.Amount,
		Description: ct.Description,
		Date:        ct.Date.Format("2006-01-02 15:04:05"),
		CreatedAt:   ct.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToCashTransactionPagination(
	cts []domain.CashTransaction,
	f filter.PaginationWithInputFilter,
	totalRows int64,
) CashTransactionResponsePagination {
	res := make([]CashTransactionResponse, len(cts))
	for i, ct := range cts {
		res[i] = toCashTransactionResponse(&ct)
	}

	return CashTransactionResponsePagination{
		Data: res,
		Meta: f.ToMeta(totalRows),
	}
}
