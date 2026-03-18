package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ModifierOption struct {
	ID              uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	ModifierGroupID uuid.UUID
	Name            string `gorm:"uniqueIndex"`
	PriceAdjustment decimal.Decimal
	CogsAdjustment  decimal.Decimal
	IsActive        bool
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	ModifierGroup *ModifierGroup `gorm:"foreignKey:ModifierGroupID"`
}
