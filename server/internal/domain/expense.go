package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Expense struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	COAID       uuid.UUID `gorm:"column:coa_id;not null"`
	Amount      decimal.Decimal
	Description string
	Date        time.Time
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	COA ChartOfAccount `gorm:"foreignKey:COAID"`
}

func ToLedgerModel(l Expense) Ledger {
	return Ledger{
		COAID:           l.COAID,
		Amount:          l.Amount,
		Notes:           &l.Description,
		ReferenceID:     l.ID,
		ReferenceType:   LedgerExpense,
		TransactionDate: l.Date,
	}
}
