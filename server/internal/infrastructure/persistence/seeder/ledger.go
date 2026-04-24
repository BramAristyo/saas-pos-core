package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedLedgerData(db *gorm.DB) {
	var coas []domain.ChartOfAccount
	db.Find(&coas)

	var shifts []domain.Shift
	db.Preload("ShiftExpenses").Find(&shifts)

	var expenses []domain.Expense
	db.Find(&expenses)

	var ledgers []domain.Ledger

	// 1. Seed Ledgers from General Expenses
	for _, exp := range expenses {
		notes := exp.Description
		ledgers = append(ledgers, domain.Ledger{
			ID:              uuid.New(),
			COAID:           exp.COAID,
			ShiftID:         nil,
			Amount:          exp.Amount,
			Notes:           &notes,
			ReferenceID:     exp.ID,
			ReferenceType:   domain.LedgerExpense,
			TransactionDate: exp.Date,
		})
	}

	if len(ledgers) > 0 {
		db.Create(&ledgers)
	}
}
