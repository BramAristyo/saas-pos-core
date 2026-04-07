package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModifierGroup struct {
	ID         uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Name       string    `gorm:"uniqueIndex"`
	IsRequired bool
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	ModifierOptions  []ModifierOption  `gorm:"foreignKey:ModifierGroupID"`
	ProductModifiers []ProductModifier `gorm:"foreignKey:ModifierGroupID"`
}
