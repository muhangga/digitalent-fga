package main

import (
	"chapter-2/session-2/internal/app"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	server := app.StartServer()

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run(":8080")
	}()
	wg.Wait()
}
