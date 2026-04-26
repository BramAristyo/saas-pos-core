package seeder

import (
	"time"
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedExpenseData(db *gorm.DB) {
	utilsCOA := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	rentCOA := uuid.MustParse("11111111-1111-1111-1111-111111111113")

	expenses := []domain.Expense{
		{
			ID:          uuid.MustParse("00000000-0000-0000-0000-000000001401"),
			COAID:       utilsCOA,
			Amount:      decimal.NewFromInt(250000),
			Description: "Electricity Bill - Jan",
			Date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			ID:          uuid.MustParse("00000000-0000-0000-0000-000000001402"),
			COAID:       utilsCOA,
			Amount:      decimal.NewFromInt(100000),
			Description: "Water Bill - Jan",
			Date:        time.Date(2024, 1, 5, 0, 0, 0, 0, time.Local),
		},
		{
			ID:          uuid.MustParse("00000000-0000-0000-0000-000000001403"),
			COAID:       rentCOA,
			Amount:      decimal.NewFromInt(5000000),
			Description: "Store Rent - Jan",
			Date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&expenses)
}
