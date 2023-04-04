package main

import (
	"fmt"
	"session-3/challange-3/internal/app"
	"session-3/challange-3/internal/config"
	"sync"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "fga-book"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=disable", host, port, user, password, dbname)

	db, err := config.OpenDB(dsn, true)
	if err != nil {
		panic(err)
	}
	defer config.CloseDB(db)

	fmt.Println("Database connection established")

	var wg sync.WaitGroup
	server := app.StartServer(db)

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run(":8080")
	}()
	wg.Wait()
}
