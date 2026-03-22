package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/config"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/persistence/database"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/persistence/seeder"
	"github.com/pressly/goose/v3"
)

// goose -dir internal/infrastructure/persistence/migrations create create_users_table sql

func main() {
	cfg := config.GetConfig()
	err := database.InitDb(cfg)
	defer database.CloseDb()

	sqlDb, _ := database.GetDb().DB()

	if len(os.Args) < 2 {
		log.Fatal("usage: migration [up|down|status|reset|version]")
	}

	if os.Args[1] == "reset-seeder" {
		err = goose.RunContext(context.Background(), "reset", sqlDb, "internal/infrastructure/persistence/migrations")
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}

		err = goose.RunContext(context.Background(), "up", sqlDb, "internal/infrastructure/persistence/migrations")
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}

		seeder.SeedAll(database.GetDb())
		fmt.Println("Migration with seeder successfully!")
		return
	}

	err = goose.RunContext(context.Background(), os.Args[1], sqlDb, "internal/infrastructure/persistence/migrations")
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration executed successfully!")
}
