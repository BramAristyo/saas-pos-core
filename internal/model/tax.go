package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Tax struct {
	ID         uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Name       string
	Percentage decimal.Decimal
	IsActive   bool
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (Tax) TableName() string {
	return "taxes"
}
