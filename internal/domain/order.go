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
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	ShiftID     uuid.UUID
	CashierID   uuid.UUID
	SalesTypeID uuid.UUID
	TaxID       *uuid.UUID
	DiscountID  *uuid.UUID
	OrderNumber string `gorm:"uniqueIndex"`

	Subtotal       decimal.Decimal
	DiscountAmount decimal.Decimal
	TaxAmount      decimal.Decimal
	ChargeAmount   decimal.Decimal
	Total          decimal.Decimal

	Status     OrderStatus
	VoidReason *string
	VoidedBy   *uuid.UUID
	VoidedAt   *time.Time
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	Shift        Shift     `gorm:"foreignKey:ShiftID"`
	Cashier      User      `gorm:"foreignKey:CashierID"`
	SalesType    SalesType `gorm:"foreignKey:SalesTypeID"`
	Tax          *Tax      `gorm:"foreignKey:TaxID"`
	Discount     *Discount `gorm:"foreignKey:DiscountID"`
	VoidedByUser *User     `gorm:"foreignKey:VoidedBy"`

	Items    []OrderItem `gorm:"foreignKey:OrderID"`
	Payments []Payment   `gorm:"foreignKey:OrderID"`
}

func (o *Order) CalculateAll() {
	// Calculate Gross Subtotal from all items
	gross := decimal.Zero
	for i := range o.Items {
		o.Items[i].CalculateSubTotal()
		gross = gross.Add(o.Items[i].Subtotal)
	}
	o.Subtotal = gross

	o.DiscountAmount = decimal.Zero
	if o.Discount != nil && o.Discount.IsActive {
		now := time.Now()
		startOk := true
		endOk := true

		if o.Discount.StartDate != nil {
			if now.Before(*o.Discount.StartDate) {
				startOk = false
			}
		}

		if o.Discount.EndDate != nil {
			if now.After(*o.Discount.EndDate) {
				endOk = false
			}
		}

		if startOk && endOk {
			switch o.Discount.Type {
			case Percentage:
				o.DiscountAmount = gross.Mul(o.Discount.Value.Div(decimal.NewFromInt(100)))
			case Fixed:
				o.DiscountAmount = o.Discount.Value
			}
		}
	}

	// Prevent discount from exceeding subtotal
	if o.DiscountAmount.GreaterThan(o.Subtotal) {
		o.DiscountAmount = o.Subtotal
	}

	// Calculate Additional Charges
	o.ChargeAmount = o.SalesType.CalculateTotalCharges(o.Subtotal)

	// Calculate Tax based (Subtotal - Discount + Charges)
	o.TaxAmount = decimal.Zero
	if o.Tax != nil {
		taxBase := o.Subtotal.Sub(o.DiscountAmount).Add(o.ChargeAmount)
		o.TaxAmount = o.Tax.CalculateTaxAmount(taxBase)
	}

	// Calculate Final Grand Total
	o.Total = o.Subtotal.Sub(o.DiscountAmount).Add(o.ChargeAmount).Add(o.TaxAmount)
}
