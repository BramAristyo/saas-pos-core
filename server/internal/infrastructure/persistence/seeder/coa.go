package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedCOAData(db *gorm.DB) {
	coas := []*domain.ChartOfAccount{
		{
			ID:       uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Name:     "Utilities",
			Type:     domain.COATypeOut,
			IsSystem: false,
		},
		{
			ID:       uuid.MustParse("11111111-1111-1111-1111-111111111112"),
			Name:     "Salary",
			Type:     domain.COATypeOut,
			IsSystem: false,
		},
		{
			ID:       uuid.MustParse("11111111-1111-1111-1111-111111111113"),
			Name:     "Rent",
			Type:     domain.COATypeOut,
			IsSystem: false,
		},
		{
			ID:       uuid.MustParse("11111111-1111-1111-1111-111111111114"),
			Name:     "Marketing",
			Type:     domain.COATypeOut,
			IsSystem: false,
		},
		{
			ID:       uuid.MustParse("11111111-1111-1111-1111-111111111115"),
			Name:     "Maintenance",
			Type:     domain.COATypeOut,
			IsSystem: false,
		},
		{
			ID:       uuid.MustParse("11111111-1111-1111-1111-111111111116"),
			Name:     "Other Income",
			Type:     domain.COATypeIn,
			IsSystem: false,
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&coas)
}
