package arrays

// Run with:
// go test -run none -bench . -benchtime 3s

// Results:
// BenchmarkLinkedListTraversal-10              729           4862667 ns/op
// BenchmarkRowTraversal-10                    1363           2607361 ns/op
// BenchmarkColTraveral-10                      991           3636270 ns/op

// Arrays allow us to allocate continguous blocks of memory
// This makes it easier to cache data as pages are read from memory
// Reading in column order, or a linked list, is not as effecient as reading in row order

import (
	"testing"
)

var fa int

func BenchmarkLinkedListTraversal(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = LinkedListTraversal()
	}

	// store output in a global variable to prevent the compiler
	// from optimizing away the function call
	fa = a
}

func BenchmarkRowTraversal(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = RowTraversal()
	}

	fa = a
}

func BenchmarkColTraveral(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = ColTraversal()
	}

	fa = a
}

func TestSemanticsDemo(t *testing.T) {
	SemanticsDemo()
}
