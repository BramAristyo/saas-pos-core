package seeder

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedModifierGroupData(db *gorm.DB) {
	milkID := uuid.New()
	sugarID := uuid.New()
	sizeID := uuid.New()

	modifierGroups := []*domain.ModifierGroup{
		{
			ID:         milkID,
			Name:       "Milk Selection",
			IsRequired: false,
			IsActive:   true,
			ModifierOptions: []domain.ModifierOption{
				{
					ModifierGroupID: milkID,
					Name:            "Oat Milk",
					PriceAdjustment: decimal.NewFromFloat(15000),
					CogsAdjustment:  decimal.NewFromFloat(10000),
					IsActive:        true,
				},
				{
					ModifierGroupID: milkID,
					Name:            "Almond Milk",
					PriceAdjustment: decimal.NewFromFloat(15000),
					CogsAdjustment:  decimal.NewFromFloat(10000),
					IsActive:        true,
				},
			},
		},
		{
			ID:         sugarID,
			Name:       "Sugar Level",
			IsRequired: true,
			IsActive:   true,
			ModifierOptions: []domain.ModifierOption{
				{
					ModifierGroupID: sugarID,
					Name:            "Normal Sugar",
					PriceAdjustment: decimal.NewFromFloat(0),
					CogsAdjustment:  decimal.NewFromFloat(0),
					IsActive:        true,
				},
				{
					ModifierGroupID: sugarID,
					Name:            "Less Sugar",
					PriceAdjustment: decimal.NewFromFloat(0),
					CogsAdjustment:  decimal.NewFromFloat(0),
					IsActive:        true,
				},
			},
		},
		{
			ID:         sizeID,
			Name:       "Cup Size",
			IsRequired: true,
			IsActive:   true,
			ModifierOptions: []domain.ModifierOption{
				{
					ModifierGroupID: sizeID,
					Name:            "Regular",
					PriceAdjustment: decimal.NewFromFloat(0),
					CogsAdjustment:  decimal.NewFromFloat(0),
					IsActive:        true,
				},
				{
					ModifierGroupID: sizeID,
					Name:            "Large",
					PriceAdjustment: decimal.NewFromFloat(5000),
					CogsAdjustment:  decimal.NewFromFloat(2000),
					IsActive:        true,
				},
			},
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&modifierGroups)
}
