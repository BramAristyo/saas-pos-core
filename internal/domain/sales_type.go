package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type SalesType struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Name      string
	IsActive  bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Charges []AdditionalCharge `gorm:"foreignKey:SalesTypeID"`
}

type AdditionalCharge struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	SalesTypeID uuid.UUID
	Name        string
	Type        AdjustmentType
	Amount      decimal.Decimal
	IsActive    bool
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (st *SalesType) CalculateTotalCharges(subtotal decimal.Decimal) decimal.Decimal {
	total := decimal.Zero
	for _, c := range st.Charges {
		if !c.IsActive {
			continue
		}

		switch c.Type {
		case Percentage:
			percent := c.Amount
			if percent.GreaterThan(decimal.NewFromInt(100)) {
				percent = decimal.NewFromInt(100)
			}

			total = total.Mul(percent.Div(decimal.NewFromInt(100)))
		case Fixed:
			total = total.Add(c.Amount)
		}
	}

	return total
}
