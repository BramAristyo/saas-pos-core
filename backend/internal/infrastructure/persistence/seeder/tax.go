package seeder

import (
	"github.com/BramAristyo/saas-pos-core/backend/internal/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedTaxData(db *gorm.DB) {
	taxes := []domain.Tax{
		{
			Name:       "TAX 11%",
			Percentage: decimal.NewFromFloat(11),
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&taxes)
}
