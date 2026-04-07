package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Shift struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	OpenedBy    uuid.UUID
	ClosedBy    *uuid.UUID
	OpeningCash decimal.Decimal
	ClosingCash *decimal.Decimal
	Notes       *string
	OpenedAt    time.Time
	ClosedAt    *time.Time
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	ShiftExpenses []ShiftExpenses `gorm:"foreignKey:ShiftID"`
	OpenedByUser  User            `gorm:"foreignKey:OpenedBy"`
	ClosedByUser  *User           `gorm:"foreignKey:ClosedBy"`
}
