package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ShiftExpensesType string

const (
	CashIn  ShiftExpensesType = "in"
	CashOut ShiftExpensesType = "out"
)

type ShiftExpenses struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	ShiftID     uuid.UUID
	Type        ShiftExpensesType
	Amount      decimal.Decimal
	Description *string
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	Shift Shift `gorm:"foreignKey:ShiftID"`
}

func (ShiftExpenses) TableName() string {
	return "shift_expenses"
}
