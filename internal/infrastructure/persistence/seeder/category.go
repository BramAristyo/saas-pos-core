package seeder

import (
	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedCategoryData(db *gorm.DB) {
	categories := []*domain.Category{
		{
			Name:        "Coffee",
			Description: "Espresso, brewed, and specialty coffee drinks",
			IsActive:    true,
		},
		{
			Name:        "Tea",
			Description: "Hot and iced teas, herbal infusions",
			IsActive:    true,
		},
		{
			Name:        "Pastries",
			Description: "Croissants, muffins, scones, and baked goods",
			IsActive:    true,
		},
		{
			Name:        "Sandwiches",
			Description: "Freshly made sandwiches and wraps",
			IsActive:    true,
		},
		{
			Name:        "Salads",
			Description: "Healthy salads and light meals",
			IsActive:    true,
		},
		{
			Name:        "Cold Drinks",
			Description: "Iced coffee, smoothies, and cold beverages",
			IsActive:    true,
		},
		{
			Name:        "Snacks",
			Description: "Cookies, chips, and small bites",
			IsActive:    true,
		},
		{
			Name:        "Breakfast",
			Description: "Breakfast sets and morning specials",
			IsActive:    true,
		},
		{
			Name:        "Desserts",
			Description: "Cakes, brownies, and sweet treats",
			IsActive:    true,
		},
		{
			Name:        "Merchandise",
			Description: "Coffee beans, mugs, and branded items",
			IsActive:    true,
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&categories)
}
