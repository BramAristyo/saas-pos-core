package main

import (
	"context"
	"log"
	"os"

	"github.com/BramAristyo/go-pos-mawish/internal/config"
	"github.com/BramAristyo/go-pos-mawish/internal/infra/persistence/database"
	"github.com/pressly/goose/v3"
)

// goose -dir internal/infra/persistence/migrations create create_users_table sql

func main() {
	cfg := config.GetConfig()
	err := database.InitDb(cfg)
	defer database.CloseDb()

	sqlDb, _ := database.GetDb().DB()

	if len(os.Args) < 2 {
		log.Fatal("usage: migration [up|down|status|reset|version]")
	}

	err = goose.RunContext(context.Background(), os.Args[1], sqlDb, "internal/infra/persistence/migrations")
	if err != nil {
		log.Fatal(err.Error())
	}
}
