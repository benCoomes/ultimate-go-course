package main

import "fmt"

// Error: cannot use 1000 (untyped int constant) as uint8 value in constant declaration (overflows)
// const myUint uint8 = 1000

func main() {

	const ui = 12345
	const uf = 3.141592

	const ti int = 12345        // type int
	const tf float64 = 3.141592 // type float64

	// ---------------------------------------------------------------------------------
	const maxInt = 9223372036854775807
	// bigger can exist because constants have 256 bits of precision
	const bigger = 9223372036854775808543522345
	// Error: constant 9223372036854775808543522345 overflows int64
	// const biggerInt int64 = 9223372036854775808543522345

	// ---------------------------------------------------------------------------------
	// Variable answer will be of type float64
	var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)
	fmt.Println(answer)

	// Constant third will be of kind floating point
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

	// Constant zero will be of kind integer
	const zero = 1 / 3 // KindInt(1) / KindInt(3)

	// This is an example of constant arithmetic between typed and
	// untyped constants. Must have like types to perform math.
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)
	// ---------------------------------------------------------------------------------

	const (
		A1 = iota // 0 : Start at 0
		B1        // 1 : Increment by 1
		C1        // 2 : Increment by 1
	)

	const (
		Ldate         = 1 << iota // 1 : Shift 1 to the left 0. 1 << 0
		Ltime                     // 2 : Shift 1 to the left 1. 1 << 1
		Lmicroseconds             // 4 : Shift 1 to the left 2. 1 << 2
		Llongfile                 // 8 : Shift 1 to the left 3. 1 << 3
		Lshortfile                // 16 : Shift 1 to the left 4. 1 << 4
		LUTC                      // 32 : Shift 1 to the left 5. 1 << 5
	)
}
