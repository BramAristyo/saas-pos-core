package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
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
		{
			Name:        "Juice",
			Description: "Freshly squeezed fruit and vegetable juices",
		},
		{
			Name:        "Smoothies",
			Description: "Blended fruit smoothies and protein shakes",
		},
		{
			Name:        "Mocktails",
			Description: "Non-alcoholic mixed drinks with unique flavors",
		},
		{
			Name:        "Hot Drinks",
			Description: "Hot chocolate, matcha, and warm beverages",
		},
		{
			Name:        "Ice Cream",
			Description: "Various flavors of ice cream and gelato",
		},
		{
			Name:        "Waffles",
			Description: "Sweet and savory waffles",
		},
		{
			Name:        "Pancakes",
			Description: "Fluffy pancakes with various toppings",
		},
		{
			Name:        "Rice Bowls",
			Description: "Rice-based meals with assorted toppings",
		},
		{
			Name:        "Pasta",
			Description: "Italian pasta dishes with rich sauces",
		},
		{
			Name:        "Pizza",
			Description: "Freshly baked pizzas with various toppings",
		},
		{
			Name:        "Burgers",
			Description: "Grilled burgers with beef, chicken, or veggie patties",
		},
		{
			Name:        "Fries",
			Description: "French fries and loaded fries",
		},
		{
			Name:        "Steak",
			Description: "Grilled steak with sides and sauces",
		},
		{
			Name:        "Soup",
			Description: "Warm soups for comfort meals",
		},
		{
			Name:        "Vegan",
			Description: "Plant-based food and beverages",
		},
		{
			Name:        "Gluten-Free",
			Description: "Gluten-free food options",
		},
		{
			Name:        "Kids Menu",
			Description: "Meals specially prepared for kids",
		},
		{
			Name:        "Seasonal",
			Description: "Limited-time seasonal menu items",
		},
		{
			Name:        "Signature",
			Description: "House signature and best-selling items",
		},
		{
			Name:        "Combos",
			Description: "Meal combos and value packages",
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&categories)
}
