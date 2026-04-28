package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShiftSchedule struct {
	ID                  uuid.UUID      `gorm:"primaryKey;default:gen_random_uuid()"`
	Name                string         `gorm:"type:varchar(50);not null"`
	StartTime           string         `gorm:"type:time;not null"`
	EndTime             string         `gorm:"type:time;not null"`
	ToleranceMinutes    int            `gorm:"not null;default:15"`
	LateIntervalMinutes int            `gorm:"not null;default:10"`
	LateDeductionAmount float64        `gorm:"type:decimal(12,2);not null;default:0"`
	CreatedAt           time.Time      `gorm:"autoCreateTime"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
