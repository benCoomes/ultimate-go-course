package main

import (
	"crypto/sha1"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func init() {
	// allocate one logical processor for scheduler to use
	//runtime.GOMAXPROCS(1)

	// or, allocate two
	runtime.GOMAXPROCS(2)
}

func main() {
	// WaitGroup is used to manage concurrency
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		//printHashes("A")
		uppercase()

		// failing to call wg.Done() will cause a deadlock! Because wg.Wait() will never unblock.
		wg.Done()
	}()

	go func() {
		//printHashes("B")
		lowercase()

		wg.Done()
	}()

	fmt.Println("Waiting to finish")
	// skipping wg.Wait() will terminate the program and the goroutines, possibly before they finish running!
	wg.Wait()

	// You can tell the scheduler to run, if you want! But this is almost always a bad idea.
	// runtime.Gosched()

	fmt.Println("\nDone")
}

func uppercase() {
	// Display alphabet 3 times
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}

func lowercase() {
	// Display alphabet 3 times
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}

func printHashes(prefix string) {
	// to see context switches: go run . | cut -c1 | grep '[AB]' | uniq

	for i := 1; i <= 50000; i++ {
		num := strconv.Itoa(i)
		sum := sha1.Sum([]byte(num))
		fmt.Printf("%s: %05d: %x\n", prefix, i, sum)
	}

	fmt.Println("Completed", prefix)
}
