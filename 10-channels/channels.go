package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	//waitForResult()
	//fanOut()
	//waitForTask()
	//pooling()
	//fanoutSemaphore()
	//fanoutBounded()
	//dropBounded()
	cancellation()
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

func waitForTask() {
	// this pattern is useful for sending work to a pool for workers

	ch := make(chan string)

	go func() {
		w := <-ch
		fmt.Println("worker : got signal : ", w)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "work"
	fmt.Println("manager :  sent signal")

	time.Sleep(time.Second)
	fmt.Println("--------------------------")
}

func pooling() {
	workers := 4
	ch := make(chan string)

	for i := 0; i < workers; i++ {
		go func(wid int) {
			for w := range ch {
				fmt.Println("worker", wid, ": got signal :", w)
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			}
			fmt.Println("worker", wid, "stopping...")
		}(i)
	}

	for i := 0; i < 20; i++ {
		ch <- fmt.Sprint("work ", i)
		fmt.Println("manager :  sent signal")
	}

	close(ch)
	fmt.Println("manager :  closed channel")

	time.Sleep(time.Second)
	fmt.Println("--------------------------")
}

func fanoutSemaphore() {
	// this pattern is helpful for scheduling a lot of work at once,
	// but limiting how much of a limited resource is used at a time (ex: database connections)
	jobs := 200
	ch := make(chan string, jobs)

	g := runtime.NumCPU()
	sem := make(chan bool, g)

	for j := 0; j < jobs; j++ {
		go func(job int) {
			sem <- true // sending on this channel will block once buffer is full
			{
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- fmt.Sprint(job)
				fmt.Println("worker: finished job :", job)
			}
			<-sem // read from channel to decrement semaphore
		}(j)
	}

	for jobs > 0 {
		j := <-ch
		jobs--
		fmt.Println("recieved job: ", j, "remaining: ", jobs)
	}
}

func fanoutBounded() {
	// when you don't want to limit then number of go routines to handle some work
	work := []string{"paper", "paper", "paper", "paper", 2000: "paper"}
	g := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for w := 0; w < g; w++ {
		go func(worker int) {
			defer wg.Done()
			for p := range ch {
				fmt.Printf("worker %d recieved signal: %s\n", worker, p)
			}
			fmt.Printf("worker %d recieved shutdown signal\n", worker)
		}(w)
	}

	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()
}

func dropBounded() {
	const cap = 100
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("worker recieved signal: ", p)
		}
	}()

	const work = 2000
	for w := 0; w < work; w++ {
		// if the channel is full, goes to the default case
		select {
		case ch <- fmt.Sprint(w):
			fmt.Println("sent work: ", w)
		default:
			fmt.Println("dropped work: ", w)
		}
	}

	close(ch)
	fmt.Println("sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------")
}

func cancellation() {
	duration := 100 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)

	// always important to call cancel at least once, otherwise memory is leaked.
	// cancel is okay to call more than once.
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "work"
	}()

	select {
	case w := <-ch:
		fmt.Println("work complete: ", w)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Millisecond * 300)
	fmt.Println("-----------------------------")
}

func cancellationBug() {
	// *************************************************************************
	// This method is similar to the one above, but has a bug - can you find it?
	// *************************************************************************

	duration := 100 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)

	// always important to call cancel at least once, otherwise memory is leaked.
	// cancel is okay to call more than once.
	defer cancel()

	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "work"
	}()

	select {
	case w := <-ch:
		fmt.Println("work complete: ", w)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Millisecond * 300)
	fmt.Println("-----------------------------")

	// *************************************************************************
	// Answer: the channel is unbuffered, so sends block until a recieve is ready.
	// In the case where <-ctx.Done() signals before <-ch, the main go routine moves on
	// and stops waiting to recieve from ch. When the other go routine tries to send on
	// ch, there will be no receivers and the goroutine will get stuck!
	// So, always make sure that work run in a goroutine doesn't get blocked sending to channels,
	// even if other goroutines have given up waiting on results from the channel.
	// *************************************************************************
}

// Important questions:
// Does the goroutine that is sending need a guarantee that its signal is recieved?
// Do we signal with data or without data?

// Channels can be in 3 states:
// open
// 0/nil
// closed
