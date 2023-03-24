package main

import "fmt"

func main() {
	fmt.Println("POINTERS ON POINTERS")

	count := 10

	println("pass by value")
	println("count - Value: ", count, "Addr: ", &count)
	increment_by_value(count)
	println("count - Value: ", count, "Addr: ", &count)

	println("pass by reference")
	println("count - Value: ", count, "Addr: ", &count)
	increment_by_reference(&count)
	println("count - Value: ", count, "Addr: ", &count)
}

func increment_by_value(i int) {
	i += 1
	println("i     - Value: ", i, "Addr: ", &i)
}

func increment_by_reference(i *int) {
	*i += 1
	println("i     - Value: ", *i, "Addr: ", i)
}
