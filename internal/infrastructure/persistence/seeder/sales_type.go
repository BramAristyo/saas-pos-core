package seeder

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedSalesTypeData(db *gorm.DB) {
	salesTypes := []*domain.SalesType{
		{
			Name:     "Dine In",
			IsActive: true,
			Charges: []domain.AdditionalCharge{
				{
					Name:     "Service Charge",
					Type:     domain.Percentage,
					Amount:   decimal.NewFromFloat(5),
					IsActive: true,
				},
				{
					Name:     "PB1",
					Type:     domain.Percentage,
					Amount:   decimal.NewFromFloat(10),
					IsActive: true,
				},
			},
		},
		{
			Name:     "Take Away",
			IsActive: true,
			Charges: []domain.AdditionalCharge{
				{
					Name:     "Packaging Fee",
					Type:     domain.Fixed,
					Amount:   decimal.NewFromFloat(2000),
					IsActive: true,
				},
				{
					Name:     "PB1",
					Type:     domain.Percentage,
					Amount:   decimal.NewFromFloat(10),
					IsActive: true,
				},
			},
		},
		{
			Name:     "Delivery",
			IsActive: true,
			Charges: []domain.AdditionalCharge{
				{
					Name:     "Delivery Fee",
					Type:     domain.Fixed,
					Amount:   decimal.NewFromFloat(5000),
					IsActive: true,
				},
				{
					Name:     "Service Fee",
					Type:     domain.Fixed,
					Amount:   decimal.NewFromFloat(1000),
					IsActive: true,
				},
				{
					Name:     "PB1",
					Type:     domain.Percentage,
					Amount:   decimal.NewFromFloat(10),
					IsActive: true,
				},
			},
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&salesTypes)
}
