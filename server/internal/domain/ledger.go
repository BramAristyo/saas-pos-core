package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ReferenceType string

const (
	LedgerShiftExpense ReferenceType = "shiftExpense"
	LedgerExpense      ReferenceType = "expense"
	LedgerSales        ReferenceType = "sales"
)

type Ledger struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	COAID           uuid.UUID
	ShiftID         uuid.UUID
	Amount          decimal.Decimal
	Notes           *string
	ReferenceID     uuid.UUID
	ReferenceType   ReferenceType
	TransactionDate time.Time
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	COA   ChartOfAccount `gorm:"foreignKey:COAID"`
	Shift *Shift         `gorm:"foreignKey:ShiftID"`
}
