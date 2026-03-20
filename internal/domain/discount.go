package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type DiscountType string

const (
	Percentage DiscountType = "percentage"
	Fixed      DiscountType = "fixed"
)

type Discount struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Name      string
	Type      DiscountType
	Value     decimal.Decimal
	StartDate *time.Time
	EndDate   *time.Time
	IsActive  bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
