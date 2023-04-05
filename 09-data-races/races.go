package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// a shared variable incremented by all goroutines
var counter int32

func main() {
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	//incrementWithRace(grs, &wg)
	//incrementAtomically(grs, &wg)
	incrementWithMutex(grs, &wg, &sync.Mutex{})

	wg.Wait()

	fmt.Printf("Counter: %d\n", counter)
}

func incrementWithRace(grs int, wg *sync.WaitGroup) {
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				value := counter
				value++
				// this debugging line changes the program's behavior!
				// it creates a data race by changing contexts between read and write.
				fmt.Println("logging - race")
				counter = value
			}

			wg.Done()
		}()
	}
}

func incrementAtomically(grs int, wg *sync.WaitGroup) {
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				atomic.AddInt32(&counter, 1)
				fmt.Println("logging - atomic")
			}
			wg.Done()
		}()
	}
}

// there is also a RWMutex which allows either: any number of readers, OR a single writer.
// That isn't helpful here though.
func incrementWithMutex(grs int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				mutex.Lock()
				{
					value := counter
					value++
					fmt.Println("logging - mutex")
					counter = value
				}
				mutex.Unlock()
			}

			wg.Done()
		}()
	}
}
