package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderItemModifier struct {
	ID               uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	OrderItemID      uuid.UUID
	ModifierOptionID *uuid.UUID
	ModifierName     string
	PriceAdjustment  decimal.Decimal
	CogsAdjustment   decimal.Decimal
	CreatedAt        time.Time `gorm:"autoCreateTime"`

	ModifierOption *ModifierOption `gorm:"foreignKey:ModifierOptionID"`
}

func (OrderItemModifier) TableName() string {
	return "order_item_modifiers"
}
