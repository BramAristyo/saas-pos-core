package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type CashType string

const (
	CashIn  CashType = "in"
	CashOut CashType = "out"
)

type CashTransaction struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	COAID       uuid.UUID
	ShiftID     *uuid.UUID
	Type        CashType
	Amount      decimal.Decimal
	Description *string
	Date        time.Time
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	COA ChartOfAccount `gorm:"foreignKey:COAID"`
}
