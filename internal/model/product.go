package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CategoryID uuid.UUID `gorm:"type:uuid"`

	Name        string          `gorm:"type:varchar(100);not null;uniqueIndex"`
	Description *string         `gorm:"type:text"`
	Price       decimal.Decimal `gorm:"type:decimal(12,2);not null;default:0"`
	Cogs        decimal.Decimal `gorm:"type:decimal(12,2);not null;default:0"`

	ImageURL *string `gorm:"type:varchar(255)"`
	IsActive bool    `gorm:"not null;default:true"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Category         *Category         `gorm:"foreignKey:CategoryID"`
	ProductModifiers []ProductModifier `gorm:"foreignKey:ProductID"`
}
