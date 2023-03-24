package main

import (
	"fmt"
	"sync"
)

func main() {

	coba := []string{"coba1", "coba2", "coba3"}
	bisa := []string{"bisa1", "bisa2", "bisa3"}

	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(8)
	for i := 1; i <= 4; i++ {
		go goroutineWg(coba, i, &wg)
		go goroutineWg(bisa, i, &wg)
	}
	wg.Wait()

	fmt.Println("=============================")

	wg.Add(8)
	for i := 1; i <= 4; i++ {
		go goroutineMutex(coba, i, &wg, &mutex)
		go goroutineMutex(bisa, i, &wg, &mutex)
	}
	wg.Wait()

}

func goroutineWg(data []string, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s %d\n", data, index)
}

func goroutineMutex(data []string, index int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	fmt.Printf("%s %d\n", data, index)
	mutex.Unlock()
}
