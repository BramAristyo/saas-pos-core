package seeder

import (
	"time"

	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func SeedExpenseData(db *gorm.DB) {
	var utilitiesCOA domain.ChartOfAccount
	db.Where("name = ?", "Utilities").First(&utilitiesCOA)

	var rentCOA domain.ChartOfAccount
	db.Where("name = ?", "Rent").First(&rentCOA)

	expenses := []domain.Expense{
		{
			COAID:       utilitiesCOA.ID,
			Amount:      decimal.NewFromFloat(250000),
			Description: "Electricity Bill - March",
			Date:        time.Now().AddDate(0, -1, 0),
		},
		{
			COAID:       utilitiesCOA.ID,
			Amount:      decimal.NewFromFloat(100000),
			Description: "Water Bill - March",
			Date:        time.Now().AddDate(0, -1, 5),
		},
		{
			COAID:       rentCOA.ID,
			Amount:      decimal.NewFromFloat(5000000),
			Description: "Store Rent - April",
			Date:        time.Now(),
		},
	}

	db.Create(&expenses)
}
