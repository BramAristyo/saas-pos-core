package main

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/infrastructure/config"
	"github.com/BramAristyo/saas-pos-core/server/internal/infrastructure/persistence/database"
	"github.com/BramAristyo/saas-pos-core/server/internal/infrastructure/persistence/seeder"
	"github.com/BramAristyo/saas-pos-core/server/pkg/logger"
)

func main() {
	cfg := config.GetConfig()
	zapLogger := logger.NewZapLogger(cfg)

	err := database.InitDb(cfg, zapLogger)
	if err != nil {
		zapLogger.Fatal(err.Error())
	}

	defer database.CloseDb()

	db := database.GetDb()
	seeder.SeedAll(db)
}
