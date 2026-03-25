package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ID           uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	OrderID      uuid.UUID
	ProductID    *uuid.UUID
	BundlingID   *uuid.UUID
	DiscountID   *uuid.UUID
	ProductName  string
	ProductPrice decimal.Decimal
	ProductCogs  decimal.Decimal
	Quantity     int

	DiscountAmount decimal.Decimal
	Subtotal       decimal.Decimal

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Product  *Product         `gorm:"foreignKey:ProductID"`
	Bundling *BundlingPackage `gorm:"foreignKey:BundlingID"`
	Discount *Discount        `gorm:"foreignKey:DiscountID"`
}

func (oi *OrderItem) CalculateSubTotal() {
	qty := decimal.NewFromInt(int64(oi.Quantity))
	price := oi.ProductPrice
	gross := price.Mul(qty)
	discount := decimal.Zero

	// Apply discount only if present, active, and within its validity period (if provided).
	if oi.Discount != nil && oi.Discount.IsActive {
		now := time.Now()
		startOk := true
		endOk := true

		if oi.Discount.StartDate != nil {
			if now.Before(*oi.Discount.StartDate) {
				startOk = false
			}
		}

		if oi.Discount.EndDate != nil {
			if now.After(*oi.Discount.EndDate) {
				endOk = false
			}
		}

		if startOk && endOk {
			switch oi.Discount.Type {
			case Percentage:
				percent := oi.Discount.Value
				if percent.GreaterThan(decimal.NewFromInt(100)) {
					percent = decimal.NewFromInt(100)
				}

				discount = gross.Mul(percent.Div(decimal.NewFromInt(100)))
			case Fixed:
				discount = oi.Discount.Value
			}
		}
	}

	if discount.GreaterThan(gross) {
		discount = gross
	}

	subtotal := gross.Sub(discount)

	oi.DiscountAmount = discount
	oi.Subtotal = subtotal
}
