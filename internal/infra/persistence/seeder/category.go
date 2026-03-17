package seeder

import (
	"github.com/BramAristyo/go-pos-mawish/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedCategoryData(db *gorm.DB) {
	categories := []*model.Category{
		{
			ID:          uuid.New(),
			Name:        "Coffee",
			Description: "Espresso, brewed, and specialty coffee drinks",
			IsActive:    true,
		},
		{
			ID:          uuid.New(),
			Name:        "Tea",
			Description: "Hot and iced teas, herbal infusions",
			IsActive:    true,
		},
		{
			ID:          uuid.New(),
			Name:        "Pastries",
			Description: "Croissants, muffins, scones, and baked goods",
			IsActive:    true,
		},
		{
			ID:          uuid.New(),
			Name:        "Sandwiches",
			Description: "Freshly made sandwiches and wraps",
			IsActive:    true,
		},
		{
			ID:          uuid.New(),
			Name:        "Salads",
			Description: "Healthy salads and light meals",
			IsActive:    true,
		},
		{
			ID:          uuid.New(),
			Name:        "Cold Drinks",
			Description: "Iced coffee, smoothies, and cold beverages",
			IsActive:    true,
		},
		{
			ID:          uuid.New(),
			Name:        "Snacks",
			Description: "Cookies, chips, and small bites",
			IsActive:    true,
		},
		{
			ID:          uuid.New(),
			Name:        "Breakfast",
			Description: "Breakfast sets and morning specials",
			IsActive:    true,
		},
		{
			ID:          uuid.New(),
			Name:        "Desserts",
			Description: "Cakes, brownies, and sweet treats",
			IsActive:    true,
		},
		{
			ID:          uuid.New(),
			Name:        "Merchandise",
			Description: "Coffee beans, mugs, and branded items",
			IsActive:    true,
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&categories)
}
