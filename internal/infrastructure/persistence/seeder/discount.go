package seeder

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedDiscountData(db *gorm.DB) {
	discounts := []domain.Discount{
		{
			Name:     "Happy Hour 10%",
			Type:     domain.Percentage,
			Value:    decimal.NewFromFloat(10),
			IsActive: true,
		},
		{
			Name:     "Fixed Discount $5",
			Type:     domain.Fixed,
			Value:    decimal.NewFromFloat(5),
			IsActive: true,
		},
		{
			Name:     "Weekend Special 15%",
			Type:     domain.Percentage,
			Value:    decimal.NewFromFloat(15),
			IsActive: true,
		},
		{
			Name:     "Employee Discount 20%",
			Type:     domain.Percentage,
			Value:    decimal.NewFromFloat(20),
			IsActive: true,
		},
		{
			Name:     "Loyalty Member $10",
			Type:     domain.Fixed,
			Value:    decimal.NewFromFloat(10),
			IsActive: true,
		},
		{
			Name:     "Student Discount 5%",
			Type:     domain.Percentage,
			Value:    decimal.NewFromFloat(5),
			IsActive: true,
		},
		{
			Name:     "New Year Promo 25%",
			Type:     domain.Percentage,
			Value:    decimal.NewFromFloat(25),
			IsActive: false,
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&discounts)
}
