package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentMethod string

const (
	Cash     PaymentMethod = "cash"
	Qris     PaymentMethod = "qris"
	Transfer PaymentMethod = "transfer"
	Other    PaymentMethod = "other"
)

type Payment struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	OrderID   uuid.UUID
	Method    PaymentMethod
	Amount    decimal.Decimal
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Order Order `gorm:"foreignKey:OrderID"`
}
