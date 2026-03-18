package domain

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	UserID      uuid.UUID
	Action      string
	Entity      string
	EntityID    *uuid.UUID
	Description *string
	IpAddress   *string

	CreatedAt time.Time `gorm:"autoCreateTime"`

	User *User `gorm:"foreignKey:UserID"`
}
