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
