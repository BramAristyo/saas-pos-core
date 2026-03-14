package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	Description string    `gorm:"type:text"`
	IsActive    bool      `gorm:"default:true;column:is_active"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	Products []Product `gorm:"foreignKey:CategoryID"`
}

// https://gorm.io/docs/conventions.html
func (Category) TableName() string {
	return "product_categories"
}
