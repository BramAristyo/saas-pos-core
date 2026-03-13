package model

import "time"

type Category struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	Description string    `gorm:"type:text"`
	IsActive    bool      `gorm:"default:true;column:is_active"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// https://gorm.io/docs/conventions.html
func (Category) TableName() string {
	return "product_categories"
}
