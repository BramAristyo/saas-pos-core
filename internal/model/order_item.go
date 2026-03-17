package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ID             uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	OrderID        uuid.UUID
	ProductID      *uuid.UUID
	BundlingID     *uuid.UUID
	DiscountID     *uuid.UUID
	ProductName    string
	ProductPrice   decimal.Decimal
	ProductCogs    decimal.Decimal
	Quantity       int
	DiscountAmount decimal.Decimal
	Subtotal       decimal.Decimal
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
