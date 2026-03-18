package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderStatus string

const (
	OrderCompleted OrderStatus = "completed"
	OrderVoided    OrderStatus = "voided"
)

type Order struct {
	ID             uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	ShiftID        uuid.UUID
	CashierID      uuid.UUID
	SalesTypeID    uuid.UUID
	TaxID          *uuid.UUID
	DiscountID     *uuid.UUID
	OrderNumber    string `gorm:"uniqueIndex"`
	Subtotal       decimal.Decimal
	DiscountAmount decimal.Decimal
	TaxAmount      decimal.Decimal
	ChargeAmount   decimal.Decimal
	Total          decimal.Decimal
	Status         OrderStatus
	VoidReason     *string
	VoidedBy       *uuid.UUID
	VoidedAt       *time.Time
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
