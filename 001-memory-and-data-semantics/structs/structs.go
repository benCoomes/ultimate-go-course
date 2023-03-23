// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"unsafe"
)

type ben struct {
	i int64
	b bool
}

type neb struct {
	i int64
	b bool
}

func main() {
	// structs are padded with bytes so that the largest data type aligns on
	// data word boundaries for the system.
	// Ex: if a system has 8 byte data words, then a struct will add padding so that all fields
	// 8 bytes or less fit into a single word.
	// Padding is also added to the end so that the entire struct ends on a word boundary.
	// This padding helps data retrival effeciency.
	// Reordering the struct fields can increase or decrease the amount of padding neccessary.
	// Still, it is better to order structs for readability unless memory use is critical.
	type extraPadding struct {
		b bool // 1
		// 7 padding bytes
		i int64   // 8
		f float32 // 4
		// 4 padding bytes
	}

	type minimalPadding struct {
		i int64   // 8
		f float32 // 4
		b bool    // 1
		// 3 byte padding
	}

	var m1 minimalPadding
	var e1 extraPadding
	e2 := extraPadding{
		b: true,
		i: 10,
		f: 3.14159,
	}

	fmt.Printf("Size of minimal padding: %v\n", unsafe.Sizeof(m1))
	fmt.Printf("Size of extra padding: %v\n", unsafe.Sizeof(e1))

	fmt.Printf("M1: %v, %v, %v\n", m1.b, m1.i, m1.f)
	fmt.Println("Flag", e2.b)
	fmt.Println("Counter", e2.i)
	fmt.Println("Pi", e2.f)

	// unnamed type with initlization
	l1 := struct {
		i int64
		b bool
	}{
		i: 1337,
		b: false,
	}
	fmt.Printf("%+v\n", l1)

	var b ben
	var n neb

	b = ben(n) // we have to explicitly cast here, even through struct defs are the same.
	fmt.Println(b.i, b.b, n.i, n.b)
}
