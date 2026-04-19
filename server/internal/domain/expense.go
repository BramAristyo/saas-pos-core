package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ExpenseCategory string

const (
	ExpenseCategoryUtilities ExpenseCategory = "utilities"
	ExpenseCategorySalary    ExpenseCategory = "salary"
	ExpenseCategoryRaw       ExpenseCategory = "raw_material"
	ExpenseCategoryRent      ExpenseCategory = "rent"
	ExpenseCategoryMarketing ExpenseCategory = "marketing"
	ExpenseCategoryOther     ExpenseCategory = "other"
)

type Expense struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	Category    ExpenseCategory
	Amount      decimal.Decimal
	Description string
	Date        time.Time
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
