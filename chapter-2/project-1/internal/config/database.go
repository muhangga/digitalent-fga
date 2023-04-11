package config

import (
	"context"
	"fmt"
	"time"

	"github.com/muhangga/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "root"
	DB_DATABASE = "fga-book"
)

func DBConn() *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_DATABASE)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Book{})

	pool, err := db.DB()
	if err != nil {
		panic("Failed to get database connection pool")
	}

	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(100)
	pool.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		panic("Failed to ping database")
	}
	return db
}

func CloseConn(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}
