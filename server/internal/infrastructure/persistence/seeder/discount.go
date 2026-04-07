package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedDiscountData(db *gorm.DB) {
	discounts := []domain.Discount{
		{
			Name:  "Happy Hour 10%",
			Type:  domain.Percentage,
			Value: decimal.NewFromFloat(10),
		},
		{
			Name:  "Fixed Discount $5",
			Type:  domain.Fixed,
			Value: decimal.NewFromFloat(5),
		},
		{
			Name:  "Weekend Special 15%",
			Type:  domain.Percentage,
			Value: decimal.NewFromFloat(15),
		},
		{
			Name:  "Employee Discount 20%",
			Type:  domain.Percentage,
			Value: decimal.NewFromFloat(20),
		},
		{
			Name:  "Loyalty Member $10",
			Type:  domain.Fixed,
			Value: decimal.NewFromFloat(10),
		},
		{
			Name:  "Student Discount 5%",
			Type:  domain.Percentage,
			Value: decimal.NewFromFloat(5),
		},
		{
			Name:  "New Year Promo 25%",
			Type:  domain.Percentage,
			Value: decimal.NewFromFloat(25),
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&discounts)
}
