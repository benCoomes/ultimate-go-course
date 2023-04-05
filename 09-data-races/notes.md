Data races can happen when a value is modified concurrently without any coordination.

Go can help detect data races with the `-race` flag:
```
go build -race
go test -race -cpu 24
``` 

The `cpu` flag sets `GOMAXPROCS`. Setting it higher creates more context switching 'chaos' to better identify races.

Then running the program produces an error:
```
$ ./data-races
logging
logging
logging
==================
WARNING: DATA RACE
Write at 0x000001217808 by goroutine 7:
  main.main.func1()
      /Users/bencoomes/repos/personal/ultimate-go-course/09-data-races/races.go:27 +0xa6

Previous write at 0x000001217808 by goroutine 6:
  main.main.func1()
      /Users/bencoomes/repos/personal/ultimate-go-course/09-data-races/races.go:27 +0xa6

Goroutine 7 (running) created at:
  main.main()
      /Users/bencoomes/repos/personal/ultimate-go-course/09-data-races/races.go:18 +0x6d

Goroutine 6 (finished) created at:
  main.main()
      /Users/bencoomes/repos/personal/ultimate-go-course/09-data-races/races.go:18 +0x6d
==================
logging
Counter: 2
Found 1 data race(s)
```