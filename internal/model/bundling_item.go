package model

import (
	"time"

	"github.com/google/uuid"
)

type BundlingItem struct {
	ID                uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	BundlingPackageID uuid.UUID `gorm:"index:idx_bundling_product"`
	ProductID         uuid.UUID `gorm:"index:idx_bundling_product"`
	Qty               int
	CreatedAt         time.Time `gorm:"autoCreateTime"`

	BundlingPackage BundlingPackage `gorm:"foreignKey:BundlingPackageID"`
	Product         Product         `gorm:"foreignKey:ProductID"`
}
