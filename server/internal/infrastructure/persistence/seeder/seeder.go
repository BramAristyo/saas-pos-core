package seeder

import (
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	SeedUserData(db)
	SeedCOAData(db)

	SeedCategoryData(db)
	SeedModifierGroupData(db)
	SeedTaxData(db)
	SeedDiscountData(db)
	SeedSalesTypeData(db)

	SeedProductData(db)
	SeedProductModifierData(db)
	SeedBundlingData(db)

	SeedEmployeeData(db)

	SeedShiftData(db)
	SeedOrderData(db)

	SeedExpenseData(db)
	SeedLedgerData(db)
}
