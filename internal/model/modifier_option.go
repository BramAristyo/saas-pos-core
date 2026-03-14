package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ModifierOption struct {
	ID              uuid.UUID       `gorm:"type:uuid;gen_random_uuid();primaryKey"`
	ModifierGroupID uuid.UUID       `gorm:"type:uuid;not null"`
	Name            string          `gorm:"type:varchar(100);not null;uniqueInde"`
	PriceAdjustment decimal.Decimal `gorm:"type:decimal(12,2);not null;default:0"`
	CogsAdjustment  decimal.Decimal `gorm:"type:decimal(12,2);not null;default:0"`
	IsActive        bool            `gorm:"not null;default:true"`
	CreatedAt       time.Time       `gorm:"autoCreateTime"`
	UpdatedAt       time.Time       `gorm:"autoUpdateTime"`

	ModifierGroup *ModifierGroup `gorm:"foreignKey:ModifierGroupID"`
}
