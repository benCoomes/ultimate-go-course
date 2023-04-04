## OS Scheduler

Concurrency is best thought of as 'out of order' execution.
Parallelism is doing multiple things at the same time.
Concurrent execution is not necessarily parallel.

'Hardware threads' - execute instructions on the hardware. Parallelism limited by number of these.
'OS threads' -  Organizes and executes a set of instructions in sequence, using a hardware thread to do the actual execution.

Three OS thread states:
1. Running - on the HW thread
2. Runnable - ready to be put on a HW thread
3. Waiting - waiting on some other resource/dependency

If only a single HW thread, how to create illusion of parallel execution?

OS threads can share a HW thread by being scheduled for execution during a limited time before switching to another thread.
There is a cost to switching context between threads, so there are only so many OS threads which can be scheduled concurrently before the context switching cost is expensive relative to the time spent actually running OS threads.

Can define a minimal time slice which sets the minium amount of time a scheduled OS thread gets to run. Still, too many threads is an issue - if there are a 1000 threads and each gets 10 ms to run, it will be 10 seconds before each thread gets to pick back up. Very laggy!

The fewer threads the OS has to schedule, the better for overall performance.

A CPU bound workload is a set of instructions that never needs to transition to a 'waiting' state. Ex: adding a set of numbers stored in memory.
An IO bound workload is a set of instructions that transitions to 'waiting' because it is awaiting results from a network call or disk read.

OS scheduler is preemptive - it needs to respond to high-priority events by stopping running threads. For us, the important fact is that scheduling is unpredictable.

For IO bound workloads, there is still a 'magic' number of OS threads where there is always 1 thread in the runnable state. But this magic number changes depending on factors outside developer control (network latency, disk speed, etc).

## Go scheduler

Goroutines are application-level threads.
They are almost the same as OS threads, except they don't have priorities.
Go scheduler is a cooperative scheduler - there are no preemptive events.
The code needs to find safe places to yield control back to the scheduler.
These safe points occur during function call transitions. 
The scheduler can run at any of these points:
* `go` keyword
* garbage collection
* system calls (async and sync)
* blocking calls

(NOTE - preemption may have been added later in go 1.12+.)

Go has a network poller.
It starts with a single thread.
Its job is to handle async system calls (networking calls).
Note - File IO is not supported async (except on Windows).

Ex: goroutine (g) wants to read from the network.
Scheduler moves g onto the network poller to wait until the read completes.
Then the scheduler can choose another goroutine to run on the OS thread.
Once network call is complete, g goes back into the local run queue with other goroutines waiting for processor time.

Ex: goroutine (g) wants to read from file IO.
Scheduler creates a new OS thread.
Moves old OS thread running g... somewhere? to wait.
Puts new goroutine on new OS thread.
When File IO is done, g goes back to LRQ.
Old OS thread is kept in waiting for future File IO.

Go scheduler can also 'steal work'. 
If a processor runs out of goroutines, it will take goroutines assigned to busier processors.

Because goroutines are so cheap, it is reasonable to have a goroutine for every web request (50,000 is okay!).

## Key Takeaway

Go swaps out goroutines when they are waiting, so that the OS thread(s) go is are always working.
From the OS perspective, go's OS threads are always busy (CPU bound).
Switching a goroutine is faster than switching the OS thread assigned to a processor, so this saves some time when context switching, compared to a multithreaded program.
It also explains why go does not need more threads than there are processors.

## Creating goroutines

Synchronization problems: when unrelated work needs access to shared resource
Orchestration problems: when different processes need to coordinate with each other

Generally, channels are for orchestration. 
Atomic operations and mutexes are for synchronization.

Code smell: Adding 1 to a wait group over and over. 
Rule: You should only create a goroutine if you know when and how it will terminate.

One helpful feature of go is that an exception will be raised if all goroutines are asleep. 
This can happen when a goroutine doesn't decrement a waitgroup, for example.
However, this won't work if there are some goroutines running on a recurring timer, since there will never be a time when all goroutines are blocked.