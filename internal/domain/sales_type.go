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
