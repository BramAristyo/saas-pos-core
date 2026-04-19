package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type COAType string

const (
	COATypeIn  COAType = "in"
	COATypeOut COAType = "out"
)

type ChartOfAccount struct {
	ID        uuid.UUID      `gorm:"primaryKey;default:gen_random_uuid()"`
	Name      string         `gorm:"size:100;not null"`
	Type      COAType        `gorm:"size:10;not null"`
	IsSystem  bool           `gorm:"default:false"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ChartOfAccount) TableName() string {
	return "chart_of_accounts"
}
