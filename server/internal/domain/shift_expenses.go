package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ShiftExpenses struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	ShiftID     uuid.UUID
	COAID       uuid.UUID `gorm:"column:coa_id;not null"`
	Amount      decimal.Decimal
	Description *string
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	Shift Shift          `gorm:"foreignKey:ShiftID"`
	COA   ChartOfAccount `gorm:"foreignKey:COAID"`
}

func (ShiftExpenses) TableName() string {
	return "shift_expenses"
}
