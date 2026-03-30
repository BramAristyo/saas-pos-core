package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Discount struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Name      string
	Type      AdjustmentType
	Value     decimal.Decimal
	StartDate *time.Time
	EndDate   *time.Time
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
