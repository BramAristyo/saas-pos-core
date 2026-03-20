package main

import (
	"log"

	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/config"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/persistence/database"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/persistence/seeder"
)

func main() {
	cfg := config.GetConfig()
	err := database.InitDb(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer database.CloseDb()

	db := database.GetDb()

	seeder.SeedUserData(db)
	seeder.SeedCategoryData(db)
	seeder.SeedProductData(db)
	seeder.SeedModifierGroupData(db)
	seeder.SeedProductModifierData(db)
	seeder.SeedTaxData(db)
	seeder.SeedDiscountData(db)
	seeder.SeedBundlingData(db)
}
