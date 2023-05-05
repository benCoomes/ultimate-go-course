// ---------------------------
// escape analysis
// ---------------------------

package main

func main() {
	escapeAnalysis()
}

// main point: the compiler can use static analysis to determine if memory should
// be allocated on the stack or the heap for variable declarations.
// Memory is allocated on the stack if the compliler can prove that it won't be
// refereced further up the stack. Otherwise, it is allocated on the heap.

// Allocating on the stack is cheaper than the heap.

// build with `go build -gcflags -m=2` to see escape analysis details

// good companion article: Understanding Allocations in Go by James Kirk
// https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d

type user struct {
	name  string
	email string
}

func escapeAnalysis() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", u2)
}

// returns a value
//
//go:noinline
func createUserV1() user {
	u := user{
		name:  "Value",
		email: "value@email.com",
	}

	println("V1", &u)

	return u
}

// returns a reference
//
//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Ref",
		email: "ref@email.com",
	}

	// sharing down the call stack allows u to be allocated on the stack
	// the createUserV2 frame will exist for the entirety of this println call,
	// so there is no chance of u being overwritten
	println("V2", &u)

	// sharing up the stack means that u must be delcared on the heap, because
	// it is possible that the createUserV2 frame is overwritten after it returns
	return &u
}
