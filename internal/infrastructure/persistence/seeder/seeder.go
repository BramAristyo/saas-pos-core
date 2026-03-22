package seeder

import (
	"sync"

	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	SeedUserData(db)

	var wg sync.WaitGroup
	wg.Add(4)
	go func() { defer wg.Done(); SeedCategoryData(db) }()
	go func() { defer wg.Done(); SeedModifierGroupData(db) }()
	go func() { defer wg.Done(); SeedTaxData(db) }()
	go func() { defer wg.Done(); SeedDiscountData(db) }()
	wg.Wait()

	SeedProductData(db)
	SeedProductModifierData(db)
	SeedBundlingData(db)

	SeedShiftData(db)
}
