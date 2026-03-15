package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type DiscountType string

const (
	Percentage DiscountType = "percentage"
	Amount     DiscountType = "amount"
)

type Discount struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Name      string
	Type      DiscountType
	Value     decimal.Decimal
	StartDate *time.Time
	EndDate   *time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
