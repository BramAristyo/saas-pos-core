package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb() error {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Singapore",
		"localhost",
		"5432",
		"postgres",
		"root",
		"go-pos",
		"disable",
	)
	var err error
	dbClient, err = gorm.Open(postgres.Open(conn))
	sqlDb, _ := dbClient.DB()

	if err = sqlDb.Ping(); err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(30 * time.Minute)

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
