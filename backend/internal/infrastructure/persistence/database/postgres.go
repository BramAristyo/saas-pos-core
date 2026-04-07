package database

import (
	"fmt"
	"time"

	"github.com/BramAristyo/saas-pos-core/backend/internal/infrastructure/config"
	"github.com/BramAristyo/saas-pos-core/backend/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config, zapLogger *logger.ZapLogger) error {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Singapore",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
		cfg.Postgres.SSLMode,
	)

	gormConfig := &gorm.Config{
		Logger: logger.NewGormZapLogger(zapLogger.GetLogger(), 200*time.Millisecond),
	}

	var err error
	dbClient, err = gorm.Open(postgres.Open(conn), gormConfig)
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()

	if err = sqlDb.Ping(); err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleCons)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenCons)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime)

	fmt.Printf("Database connected successfully")
	return nil
}

func CloseDb() {
	conn, _ := dbClient.DB()
	conn.Close()
}

func GetDb() *gorm.DB {
	return dbClient
}
