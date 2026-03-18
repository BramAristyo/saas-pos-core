package domain

import (
	"time"

	"github.com/google/uuid"
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
	IsActive  bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	OpenShifts   []Shift `gorm:"foreignKey:OpenedBy"`
	ClosedShifts []Shift `gorm:"foreignKey:ClosedBy"`
}
