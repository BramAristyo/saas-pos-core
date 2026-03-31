package seeder

import (
	"math/rand"

	"github.com/BramAristyo/go-pos-mawish/internal/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedProductData(db *gorm.DB) {
	var categories []domain.Category
	db.Find(&categories)

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
		decimal.NewFromInt(20000), decimal.NewFromInt(28000), decimal.NewFromInt(22000), decimal.NewFromInt(18000), decimal.NewFromInt(45000),
		decimal.NewFromInt(32000), decimal.NewFromInt(22000), decimal.NewFromInt(38000), decimal.NewFromInt(35000), decimal.NewFromInt(30000),
		decimal.NewFromInt(25000), decimal.NewFromInt(25000), decimal.NewFromInt(28000), decimal.NewFromInt(22000), decimal.NewFromInt(15000),
		decimal.NewFromInt(32000), decimal.NewFromInt(55000), decimal.NewFromInt(48000), decimal.NewFromInt(42000), decimal.NewFromInt(45000),
		decimal.NewFromInt(65000), decimal.NewFromInt(85000), decimal.NewFromInt(95000), decimal.NewFromInt(88000), decimal.NewFromInt(20000),
		decimal.NewFromInt(35000), decimal.NewFromInt(30000), decimal.NewFromInt(18000), decimal.NewFromInt(45000), decimal.NewFromInt(42000),
	}

	productCogs := []decimal.Decimal{
		decimal.NewFromInt(8000), decimal.NewFromInt(12000), decimal.NewFromInt(9000), decimal.NewFromInt(7000), decimal.NewFromInt(18000),
		decimal.NewFromInt(14000), decimal.NewFromInt(9000), decimal.NewFromInt(16000), decimal.NewFromInt(15000), decimal.NewFromInt(12000),
		decimal.NewFromInt(10000), decimal.NewFromInt(10000), decimal.NewFromInt(12000), decimal.NewFromInt(9000), decimal.NewFromInt(6000),
		decimal.NewFromInt(14000), decimal.NewFromInt(22000), decimal.NewFromInt(20000), decimal.NewFromInt(18000), decimal.NewFromInt(20000),
		decimal.NewFromInt(28000), decimal.NewFromInt(40000), decimal.NewFromInt(45000), decimal.NewFromInt(42000), decimal.NewFromInt(8000),
		decimal.NewFromInt(15000), decimal.NewFromInt(12000), decimal.NewFromInt(7000), decimal.NewFromInt(20000), decimal.NewFromInt(18000),
	}

	products := make([]domain.Product, len(productNames))
	for i := range productNames {
		category := categories[rand.Intn(len(categories))]
		products[i] = domain.Product{
			Name:        productNames[i],
			Description: &productDescriptions[i],
			Price:       productPrices[i],
			Cogs:        productCogs[i],
			CategoryID:  category.ID,
		}
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&products)
}
