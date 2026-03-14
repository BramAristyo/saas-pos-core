package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductModifier struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProductID       uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_product_modifier;not null"`
	ModifierGroupID uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_product_modifier;not null"`

	Product       Product       `gorm:"foreignKey:ProductID"`
	ModifierGroup ModifierGroup `gorm:"foreignKey:ModifierGroupID"`
	CreatedAt     time.Time     `gorm:"autoCreateTime"`
}
