package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Name        string    `gorm:"uniqueIndex"`
	Description string
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Products []Product `gorm:"foreignKey:CategoryID"`
}

// https://gorm.io/docs/conventions.html
func (Category) TableName() string {
	return "product_categories"
}
