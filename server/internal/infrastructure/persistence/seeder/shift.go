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

	shifts := []domain.Shift{
		{
			OpenedBy:    openedBy,
			ClosedBy:    closedBy,
			OpeningCash: openingCash,
			ClosingCash: &closingCash,
			Notes:       &notes,
			OpenedAt:    time.Now().Add(-8 * time.Hour),
			ClosedAt:    &closedAt,
		},
		{
			OpenedBy:    openedBy,
			OpeningCash: decimal.NewFromFloat(300000),
			OpenedAt:    time.Now().Add(-2 * time.Hour),
		},
	}

	db.Create(&shifts)
}
