package db

import (
	"fmt"
	"log"

	"go_home-main/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenPostgres は DSN で Gorm(Postgres) を開いてマイグレーションまで行う簡易ヘルパー
// DSN 例: "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Tokyo"
func OpenPostgres(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	// マイグレーション（必要に応じて実行）
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return nil, fmt.Errorf("auto migrate: %w", err)
	}

	log.Println("connected to postgres and migrated")
	return db, nil
}
