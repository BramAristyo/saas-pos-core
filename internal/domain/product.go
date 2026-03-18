package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID         uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	CategoryID uuid.UUID

	Name        string `gorm:"uniqueIndex"`
	Description *string
	Price       decimal.Decimal
	Cogs        decimal.Decimal

	ImageURL *string
	IsActive bool

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Category         *Category         `gorm:"foreignKey:CategoryID"`
	ProductModifiers []ProductModifier `gorm:"foreignKey:ProductID"`
	BundlingItems    []BundlingItem    `gorm:"foreignKey:ProductID"`
}
