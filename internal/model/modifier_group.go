package model

import (
	"time"

	"github.com/google/uuid"
)

type ModifierGroup struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	IsRequired bool      `gorm:"not null;default:false"`
	IsActive   bool      `gorm:"not null;default:true"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	ProductModifiers []ProductModifier `gorm:"foreignKey:ModifierGroupID"`
}
