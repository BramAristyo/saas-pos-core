package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BundlingPackage struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Name        string
	Description *string
	Price       decimal.Decimal
	Cogs        decimal.Decimal
	ImageURL    *string
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	BundlingItems []BundlingItem `gorm:"foreignKey:BundlingPackageID"`
}
