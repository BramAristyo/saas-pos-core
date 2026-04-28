package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Code       string         `gorm:"type:varchar(100);not null;unique"`
	Name       string         `gorm:"type:varchar(100);not null"`
	Phone      *string        `gorm:"type:varchar(20)"`
	BaseSalary float64        `gorm:"type:decimal(12,2);not null;default:0"`
	PinHash       string         `gorm:"type:varchar(255);not null"`
	HasChangedPIN bool           `gorm:"default:false"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
