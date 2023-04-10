package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//waitForResult()
	fanOut()
}

func waitForResult() {
	ch := make(chan string) // unbuffered channel that signals with string data

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("worker: sent signal")
	}()

	p := <-ch // blocking receive
	// worker and main thread are running in parallel, so no guaruntee about the order of println
	// unless GOMAXPROCS=1, then worker println will always run first.
	fmt.Println("manager: recieved signal - ", p)

	time.Sleep(time.Duration(100) * time.Millisecond)
	fmt.Println("------------------------------")
}

func fanOut() {
	// fan out patterns are dangerous because they can create a lot of load
	// for system running go program and/or external systems
	works := 2000
	ch := make(chan string, works)

	for w := 0; w < works; w++ {
		go func(work int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("worker : sent signal :", work)
		}(w)
	}

	for works > 0 {
		w := <-ch
		works--
		fmt.Println(w)
		fmt.Println("manager : got signal : ", works)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------")
}

// Important questions:
// Does the goroutine that is sending need a guarantee that its signal is recieved?
// Do we signal with data or without data?

// Channels can be in 3 states:
// open
// 0/nil
// closed
