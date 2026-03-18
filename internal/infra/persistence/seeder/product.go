package seeder

import (
	"math/rand"

	"github.com/BramAristyo/go-pos-mawish/internal/model"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedProductData(db *gorm.DB) {
	var categories []model.Category
	db.Where("is_active = ?", true).Find(&categories)

	if len(categories) == 0 {
		return
	}

	productNames := []string{
		"Espresso", "Latte", "Green Tea", "Croissant", "Turkey Sandwich",
		"Cappuccino", "Americano", "Mocha", "Flat White", "Cold Brew",
		"Earl Grey Tea", "Chamomile Tea", "Iced Peach Tea", "Blueberry Muffin", "Chocolate Chip Cookie",
		"Bagel with Cream Cheese", "Avocado Toast", "Caesar Salad", "Grilled Cheese", "Veggie Wrap",
		"Club Sandwich", "Beef Burger", "Margherita Pizza", "Pasta Carbonara", "Lemonade",
		"Smoothies", "Hot Chocolate", "Scone", "Quiche", "Chicken Caesar Wrap",
	}
	productDescriptions := []string{
		"Strong coffee", "Milk coffee", "Refreshing tea", "Buttery pastry", "Hearty sandwich",
		"Rich and creamy coffee", "Classic black coffee", "Chocolate-flavored latte", "Velvety smooth coffee", "Slow-steeped iced coffee",
		"Bergamot-infused black tea", "Calming herbal tea", "Sweet and refreshing", "Freshly baked with real berries", "Gooey and sweet",
		"Toasted and savory", "Healthy and delicious", "Fresh greens and parmesan", "Comforting and melty", "Nutritious and flavorful",
		"Classic triple-decker", "Juicy beef patty", "Simple and authentic", "Creamy and savory", "Zesty and tart",
		"Mixed berry blend", "Rich and warming", "Classic English pastry", "Savory custard tart", "Fresh and filling",
	}
	productPrices := []decimal.Decimal{
		decimal.NewFromFloat(2.5), decimal.NewFromFloat(3.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(1.5), decimal.NewFromFloat(4.0),
		decimal.NewFromFloat(3.5), decimal.NewFromFloat(2.0), decimal.NewFromFloat(4.5), decimal.NewFromFloat(3.8), decimal.NewFromFloat(4.0),
		decimal.NewFromFloat(2.5), decimal.NewFromFloat(2.5), decimal.NewFromFloat(3.0), decimal.NewFromFloat(2.5), decimal.NewFromFloat(1.8),
		decimal.NewFromFloat(3.2), decimal.NewFromFloat(6.5), decimal.NewFromFloat(7.0), decimal.NewFromFloat(5.5), decimal.NewFromFloat(6.0),
		decimal.NewFromFloat(8.5), decimal.NewFromFloat(9.5), decimal.NewFromFloat(10.0), decimal.NewFromFloat(11.0), decimal.NewFromFloat(2.8),
		decimal.NewFromFloat(4.5), decimal.NewFromFloat(3.5), decimal.NewFromFloat(2.2), decimal.NewFromFloat(5.0), decimal.NewFromFloat(7.5),
	}

	products := make([]model.Product, len(productNames))
	for i := range productNames {
		category := categories[rand.Intn(len(categories))]
		products[i] = model.Product{
			Name:        productNames[i],
			Description: &productDescriptions[i],
			Price:       productPrices[i],
			CategoryID:  category.ID,
			IsActive:    true,
		}
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&products)
}
