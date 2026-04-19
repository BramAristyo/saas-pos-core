package seeder

import (
	"time"

	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func SeedShiftData(db *gorm.DB) {
	var users []domain.User
	db.Limit(2).Find(&users)

	if len(users) < 1 {
		return
	}

	var utilitiesCOA domain.ChartOfAccount
	db.Where("name = ?", "Utilities").First(&utilitiesCOA)

	var incomeCOA domain.ChartOfAccount
	db.Where("name = ?", "Other Income").First(&incomeCOA)

	openedBy := users[0].ID
	var closedBy *uuid.UUID
	if len(users) > 1 {
		closedBy = &users[1].ID
	}

	closedAt := time.Now()
	openingCash := decimal.NewFromFloat(500000)
	closingCash := decimal.NewFromFloat(1500000)
	notes := "Shift pagi"

	desc1 := "Beli bahan baku"
	desc2 := "Tip dari pelanggan"
	desc3 := "Beli plastik"

	shifts := []domain.Shift{
		{
			OpenedBy:    openedBy,
			ClosedBy:    closedBy,
			OpeningCash: openingCash,
			ClosingCash: &closingCash,
			Notes:       &notes,
			OpenedAt:    time.Now().Add(-8 * time.Hour),
			ClosedAt:    &closedAt,
			ShiftExpenses: []domain.ShiftExpenses{
				{
					COAID:       utilitiesCOA.ID,
					Amount:      decimal.NewFromFloat(50000),
					Description: &desc1,
				},
				{
					COAID:       incomeCOA.ID,
					Amount:      decimal.NewFromFloat(20000),
					Description: &desc2,
				},
				{
					COAID:       utilitiesCOA.ID,
					Amount:      decimal.NewFromFloat(15000),
					Description: &desc3,
				},
			},
		},
		{
			OpenedBy:    openedBy,
			OpeningCash: decimal.NewFromFloat(300000),
			OpenedAt:    time.Now().Add(-2 * time.Hour),
			ShiftExpenses: []domain.ShiftExpenses{
				{
					COAID:       utilitiesCOA.ID,
					Amount:      decimal.NewFromFloat(30000),
					Description: &desc1,
				},
				{
					COAID:       incomeCOA.ID,
					Amount:      decimal.NewFromFloat(10000),
					Description: &desc2,
				},
			},
		},
	}

	db.Create(&shifts)
}
