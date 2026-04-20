package dto

import (
	"time"

	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type LedgerResponse struct {
	ID              uuid.UUID            `json:"id"`
	COAName         string               `json:"coaName"`
	COAType         domain.COAType       `json:"coaType"`
	Amount          decimal.Decimal      `json:"amount"`
	ReferenceID     uuid.UUID            `json:"referenceID"`
	ReferenceType   domain.ReferenceType `json:"referenceType"`
	TransactionDate time.Time            `json:"transactionDate"`
	CreatedAt       time.Time            `json:"createdAt"`
}

type TransactionSummaryResponse struct {
	TransactionDate time.Time            `json:"transactionDate"`
	ReferenceType   domain.ReferenceType `json:"referenceType"`
	COAName         string               `json:"coaName"`
	COAType         domain.COAType       `json:"coaType"`
	
}
