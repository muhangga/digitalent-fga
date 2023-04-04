package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func OpenDB(dsn string, setLimits bool) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if setLimits {
		fmt.Println("Setting database limits")
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database successfully connect")

	return db, nil
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
