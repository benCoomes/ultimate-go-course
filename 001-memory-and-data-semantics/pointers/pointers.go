package main

import "fmt"

func main() {
	fmt.Println("POINTERS ON POINTERS")

	// pass by value
	count := 10
	println("count - Value: ", count, "Addr: ", &count)
	increment_by_value(count)
	println("count - Value: ", count, "Addr: ", &count)
}

func increment_by_value(i int) {
	i += 1
	println("i     - Value: ", i, "Addr: ", &i)
}
