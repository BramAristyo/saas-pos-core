package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedCOAData(db *gorm.DB) {
	coas := []*domain.ChartOfAccount{
		{
			Name:     "Utilities",
			Type:     domain.COATypeOut,
			IsSystem: false,
		},
		{
			Name:     "Salary",
			Type:     domain.COATypeOut,
			IsSystem: false,
		},
		{
			Name:     "Rent",
			Type:     domain.COATypeOut,
			IsSystem: false,
		},
		{
			Name:     "Marketing",
			Type:     domain.COATypeOut,
			IsSystem: false,
		},
		{
			Name:     "Maintenance",
			Type:     domain.COATypeOut,
			IsSystem: false,
		},
		{
			Name:     "Other Income",
			Type:     domain.COATypeIn,
			IsSystem: false,
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&coas)
}
