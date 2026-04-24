package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedModifierGroupData(db *gorm.DB) {
	milkID := uuid.MustParse("d57bb96c-433e-4819-911b-3caedf2f42b3")
	sugarID := uuid.MustParse("5a0b379d-bca3-493a-a216-81f6b1ba9965")
	sizeID := uuid.MustParse("18a6ee9f-75ed-4af9-b36e-dc5a9e1b14ed")
	iceID := uuid.MustParse("4fdda7fb-007d-489f-b7ab-ecacf8e935ab")
	toppingID := uuid.MustParse("00a1e5ec-5780-447a-918d-9d18aba4d1b5")
	beanID := uuid.MustParse("7363b6b7-fb4e-4356-b86d-086312e07919")
	donenessID := uuid.MustParse("fc607754-f809-4712-81d3-fc639a21f4d6")
	sideID := uuid.MustParse("deda13b7-1407-4b3d-8243-d03935a8fc46")

	modifierGroups := []*domain.ModifierGroup{
		{
			ID:         milkID,
			Name:       "Milk Selection",
			IsRequired: false,
			ModifierOptions: []domain.ModifierOption{
				{ModifierGroupID: milkID, Name: "Full Cream Milk", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
				{ModifierGroupID: milkID, Name: "Oat Milk", PriceAdjustment: decimal.NewFromFloat(15000), CogsAdjustment: decimal.NewFromFloat(10000)},
				{ModifierGroupID: milkID, Name: "Almond Milk", PriceAdjustment: decimal.NewFromFloat(15000), CogsAdjustment: decimal.NewFromFloat(10000)},
				{ModifierGroupID: milkID, Name: "Soy Milk", PriceAdjustment: decimal.NewFromFloat(12000), CogsAdjustment: decimal.NewFromFloat(8000)},
			},
		},
		{
			ID:         sugarID,
			Name:       "Sugar Level",
			IsRequired: true,
			ModifierOptions: []domain.ModifierOption{
				{ModifierGroupID: sugarID, Name: "Normal Sugar", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
				{ModifierGroupID: sugarID, Name: "Less Sugar (50%)", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
				{ModifierGroupID: sugarID, Name: "No Sugar", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
				{ModifierGroupID: sugarID, Name: "Extra Sugar", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
			},
		},
		{
			ID:         sizeID,
			Name:       "Cup Size",
			IsRequired: true,
			ModifierOptions: []domain.ModifierOption{
				{ModifierGroupID: sizeID, Name: "Regular", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
				{ModifierGroupID: sizeID, Name: "Large", PriceAdjustment: decimal.NewFromFloat(5000), CogsAdjustment: decimal.NewFromFloat(2000)},
				{ModifierGroupID: sizeID, Name: "Extra Large", PriceAdjustment: decimal.NewFromFloat(10000), CogsAdjustment: decimal.NewFromFloat(4000)},
			},
		},
		{
			ID:         iceID,
			Name:       "Ice Level",
			IsRequired: true,
			ModifierOptions: []domain.ModifierOption{
				{ModifierGroupID: iceID, Name: "Normal Ice", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
				{ModifierGroupID: iceID, Name: "Less Ice", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
				{ModifierGroupID: iceID, Name: "No Ice", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
			},
		},
		{
			ID:         toppingID,
			Name:       "Extra Toppings",
			IsRequired: false,
			ModifierOptions: []domain.ModifierOption{
				{ModifierGroupID: toppingID, Name: "Pearl (Boba)", PriceAdjustment: decimal.NewFromFloat(5000), CogsAdjustment: decimal.NewFromFloat(2000)},
				{ModifierGroupID: toppingID, Name: "Grass Jelly", PriceAdjustment: decimal.NewFromFloat(4000), CogsAdjustment: decimal.NewFromFloat(1500)},
				{ModifierGroupID: toppingID, Name: "Whipped Cream", PriceAdjustment: decimal.NewFromFloat(6000), CogsAdjustment: decimal.NewFromFloat(2500)},
				{ModifierGroupID: toppingID, Name: "Caramel Drizzle", PriceAdjustment: decimal.NewFromFloat(3000), CogsAdjustment: decimal.NewFromFloat(1000)},
			},
		},
		{
			ID:         beanID,
			Name:       "Coffee Bean Type",
			IsRequired: true,
			ModifierOptions: []domain.ModifierOption{
				{ModifierGroupID: beanID, Name: "House Blend", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
				{ModifierGroupID: beanID, Name: "100% Arabica", PriceAdjustment: decimal.NewFromFloat(5000), CogsAdjustment: decimal.NewFromFloat(2000)},
				{ModifierGroupID: beanID, Name: "Single Origin", PriceAdjustment: decimal.NewFromFloat(8000), CogsAdjustment: decimal.NewFromFloat(3000)},
			},
		},
		{
			ID:         donenessID,
			Name:       "Meat Doneness",
			IsRequired: true,
			ModifierOptions: []domain.ModifierOption{
				{ModifierGroupID: donenessID, Name: "Rare", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
				{ModifierGroupID: donenessID, Name: "Medium Rare", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
				{ModifierGroupID: donenessID, Name: "Medium", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
				{ModifierGroupID: donenessID, Name: "Well Done", PriceAdjustment: decimal.NewFromFloat(0), CogsAdjustment: decimal.NewFromFloat(0)},
			},
		},
		{
			ID:         sideID,
			Name:       "Side Dish",
			IsRequired: false,
			ModifierOptions: []domain.ModifierOption{
				{ModifierGroupID: sideID, Name: "French Fries", PriceAdjustment: decimal.NewFromFloat(15000), CogsAdjustment: decimal.NewFromFloat(5000)},
				{ModifierGroupID: sideID, Name: "Side Salad", PriceAdjustment: decimal.NewFromFloat(12000), CogsAdjustment: decimal.NewFromFloat(4000)},
				{ModifierGroupID: sideID, Name: "Mashed Potato", PriceAdjustment: decimal.NewFromFloat(18000), CogsAdjustment: decimal.NewFromFloat(6000)},
			},
		},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&modifierGroups)
}
