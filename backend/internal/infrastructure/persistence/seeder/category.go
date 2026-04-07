package seeder

import (
	"github.com/BramAristyo/saas-pos-core/backend/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedCategoryData(db *gorm.DB) {
	categories := []*domain.Category{
		{
			Name:        "Coffee",
			Description: "Espresso, brewed, and specialty coffee drinks",
		},
		{
			Name:        "Tea",
			Description: "Hot and iced teas, herbal infusions",
		},
		{
			Name:        "Pastries",
			Description: "Croissants, muffins, scones, and baked goods",
		},
		{
			Name:        "Sandwiches",
			Description: "Freshly made sandwiches and wraps",
		},
		{
			Name:        "Salads",
			Description: "Healthy salads and light meals",
		},
		{
			Name:        "Cold Drinks",
			Description: "Iced coffee, smoothies, and cold beverages",
		},
		{
			Name:        "Snacks",
			Description: "Cookies, chips, and small bites",
		},
		{
			Name:        "Breakfast",
			Description: "Breakfast sets and morning specials",
		},
		{
			Name:        "Desserts",
			Description: "Cakes, brownies, and sweet treats",
		},
		{
			Name:        "Merchandise",
			Description: "Coffee beans, mugs, and branded items",
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&categories)
}
