// package main

// import (
// 	"fmt"
// 	"sync"
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// )

// func goroutineAcak(wg *sync.WaitGroup) {

// 	wg.Add(8)
// 	for i := 1; i <= 4; i++ {
// 		var id int
// 		var name string

// 		if id%2 == 0 {
// 			name = "bisa"
// 		} else {
// 			name = "coba"
// 		}
// 		items := []string{name + "1", name + "2", name + "3"}
// 		fmt.Printf("%s %d\n", items, id)
// 	}
// 	wg.Wait()
// 	wg.Done()

// }

// func goroutineRapih(wg *sync.WaitGroup, mutex *sync.Mutex) {

// 	wg.Add(8)
// 	for i := 1; i <= 4; i++ {
// 		var id int
// 		var name string

// 		if id%2 == 0 {
// 			name = "bisa"
// 		} else {
// 			name = "coba"
// 		}
// 		items := []string{name + "1", name + "2", name + "3"}
// 		mutex.Lock()
// 		fmt.Printf("%s %d\n", items, id)
// 		mutex.Unlock()
// 	}
// 	wg.Wait()
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	var mutex sync.Mutex

// 	fmt.Println("Goroutine Acak")
// 	goroutineAcak(&wg)

// 	fmt.Println("Goroutine Rapih")
// 	goroutineRapih(&wg, &mutex)

// 	wg.Wait()

// }

//func processInterface1(data interface{}, wg *sync.WaitGroup, mutex *sync.Mutex) {
//	defer wg.Done()
//	for i := 1; i <= 4; i++ {
//		mutex.Lock()
//		fmt.Println("Interface1:", data)
//		mutex.Unlock()
//	}
//}
//
//func processInterface2(data interface{}, wg *sync.WaitGroup, mutex *sync.Mutex) {
//	defer wg.Done()
//	for i := 1; i <= 4; i++ {
//		mutex.Lock()
//		fmt.Println("Interface2:", data)
//		mutex.Unlock()
//	}
//}
//
//func main() {
//	data1 := "Data1"
//	data2 := 42
//
//	var wg sync.WaitGroup
//	var mutex sync.Mutex
//
//	for i := 1; i <= 4; i++ {
//		wg.Add(2)
//		go processInterface1(data1, &wg, &mutex)
//		go processInterface2(data2, &wg, &mutex)
//	}
//
//	wg.Wait()
//}

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
