package seeder

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func SeedBundlingData(db *gorm.DB) {
	var products []domain.Product
	db.Limit(5).Find(&products)

	if len(products) < 2 {
		return
	}

	description1 := "Perfect morning combo"
	bundlingPackage := domain.BundlingPackage{
		Name:        "Breakfast Bundle",
		Description: &description1,
		Price:       decimal.NewFromInt(35000),
		Cogs:        decimal.NewFromInt(15000),
	}

	if err := db.Where(domain.BundlingPackage{Name: "Breakfast Bundle"}).FirstOrCreate(&bundlingPackage).Error; err != nil {
		return
	}

	bundlingItems := []domain.BundlingItem{
		{
			BundlingPackageID: bundlingPackage.ID,
			ProductID:         products[0].ID,
			Qty:               1,
		},
		{
			BundlingPackageID: bundlingPackage.ID,
			ProductID:         products[1].ID,
			Qty:               1,
		},
	}

	for _, item := range bundlingItems {
		db.Where(domain.BundlingItem{
			BundlingPackageID: item.BundlingPackageID,
			ProductID:         item.ProductID,
		}).FirstOrCreate(&item)
	}
}
