package model

import (
	"time"

	"github.com/google/uuid"
)

type AdditionalType struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Name      string
	IsActive  bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
