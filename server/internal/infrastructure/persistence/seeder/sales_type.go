package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedSalesTypeData(db *gorm.DB) {
	salesTypes := []*domain.SalesType{
		{
			Name: "Dine In",
			Charges: []domain.AdditionalCharge{
				{
					Name:   "Service Charge",
					Type:   domain.Percentage,
					Amount: decimal.NewFromFloat(5),
				},
				{
					Name:   "PB1",
					Type:   domain.Percentage,
					Amount: decimal.NewFromFloat(10),
				},
			},
		},
		{
			Name: "Take Away",
			Charges: []domain.AdditionalCharge{
				{
					Name:   "Packaging Fee",
					Type:   domain.Fixed,
					Amount: decimal.NewFromFloat(2000),
				},
				{
					Name:   "PB1",
					Type:   domain.Percentage,
					Amount: decimal.NewFromFloat(10),
				},
			},
		},
		{
			Name: "Delivery",
			Charges: []domain.AdditionalCharge{
				{
					Name:   "Delivery Fee",
					Type:   domain.Fixed,
					Amount: decimal.NewFromFloat(5000),
				},
				{
					Name:   "Service Fee",
					Type:   domain.Fixed,
					Amount: decimal.NewFromFloat(1000),
				},
				{
					Name:   "PB1",
					Type:   domain.Percentage,
					Amount: decimal.NewFromFloat(10),
				},
			},
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&salesTypes)
}
