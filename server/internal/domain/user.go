package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleCashier Role = "cashier"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Role      Role
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
