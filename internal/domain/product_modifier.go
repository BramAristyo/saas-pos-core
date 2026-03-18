package domain

import (
	"time"

	"github.com/google/uuid"
)

type ProductModifier struct {
	ID              uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	ProductID       uuid.UUID `gorm:"uniqueIndex:idx_product_modifier"`
	ModifierGroupID uuid.UUID `gorm:"uniqueIndex:idx_product_modifier"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`

	Product       Product       `gorm:"foreignKey:ProductID"`
	ModifierGroup ModifierGroup `gorm:"foreignKey:ModifierGroupID"`
}
