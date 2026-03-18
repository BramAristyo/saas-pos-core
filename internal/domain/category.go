package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Name        string    `gorm:"uniqueIndex"`
	Description string
	IsActive    bool      `gorm:"column:is_active"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	Products []Product `gorm:"foreignKey:CategoryID"`
}

// https://gorm.io/docs/conventions.html
func (Category) TableName() string {
	return "product_categories"
}
