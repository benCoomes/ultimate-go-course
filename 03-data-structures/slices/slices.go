package slices

import (
	"fmt"
	"unsafe"
)

// https://go.dev/blog/slices-intro

// Slice
// ptr - pointer to head of backing array
// len - number of elements that can be read/written
// cap - number of elements in backing array (room for future growth)

func SlicesDemo() {
	fruits := make([]string, 5)
	fruits[0] = "Apfel"
	fruits[1] = "Bannana"
	fruits[2] = "Cantalope"
	fruits[3] = "Durian"
	fruits[4] = "Elderberry"

	//fruits[5] = "Runtime error"

	inspectSlice(fruits)

	// declare a nil slice of strings
	var nilSlice []string
	// Note: nice pattern to use var for zero values only, makes identification easy
	// var nilSlice2 []string = []string{} // Ex: use ':=' here instead
	fmt.Printf("nilSlice is nil? %v\n", nilSlice == nil)

	emptySlice := []string{}
	fmt.Printf("emptySlice is nil? %v\n", emptySlice == nil)

	fmt.Printf("%v\n", unsafe.Sizeof(emptySlice))
	fmt.Printf("fruits:      %p\n", fruits)
	fmt.Printf("&fruits:     %p\n", &fruits)
	fmt.Printf("&fruits[1]:  %p\n", &fruits[1])
	fmt.Printf("&struct{}{}: %p\n", &struct{}{})
	fmt.Printf("nilSlice:    %p\n", nilSlice)
	fmt.Printf("emptySlice:  %p\n", emptySlice) // empty slice points to the empty struct!
}

func AppendDemo() {
	data := make([]string, 0)

	for i := 0; i < 1e5; i++ {
		val := fmt.Sprintf("value %d", i)
		lastCap := cap(data)
		data = append(data, val)
		newCap := cap(data)

		if lastCap != newCap && lastCap != 0 {
			fmt.Printf("Old Cap: %d, New Cap: %d (%f%% increase)\n", lastCap, newCap, float64(newCap-lastCap)/float64(lastCap)*100)
		}
	}
}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d], Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d], %p, %s\n", i, &slice[i], s)
	}
}

/*

Memory leaks in go

Leaks in go are when memory on the heap is referenced, but no longer needed by the program.
For example, if a goroutine is blocked and has a reference to memory, the memory is 'leaked' because it cannot be freed.

The common places leaks happen:
1. blocked/long-running goroutines
2. maps that have k/v pairs which are no longer needed (ex: cache that never evicts values)
3. append calls, if the value passed in is not replaced by the returned value. This keeps old backing arrays alive because they still are referenced.
4. APIs with a 'close' function, where close might not be called

*/
