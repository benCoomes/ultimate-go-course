package arrays

import "fmt"

const (
	rows = 2 * 1024
	cols = 2 * 1024
)

var matrix [rows][cols]byte

type data struct {
	v byte
	p *data
}

var list *data

func init() {
	var last *data

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			var d data
			if list == nil {
				list = &d
			}
			if last != nil {
				last.p = &d
			}
			last = &d

			if row%2 == 0 {
				matrix[row][col] = 0xFF
				d.v = 0xFF
			}
		}
	}
}

func LinkedListTraversal() int {
	count := 0
	d := list
	for d != nil {
		if d.v == 0xFF {
			count++
		}
		d = d.p
	}

	return count
}

func RowTraversal() int {
	count := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				count++
			}
		}
	}
	return count
}

func ColTraversal() int {
	count := 0
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				count++
			}
		}
	}
	return count
}

func SemanticsDemo() {
	var fruits [5]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "orange"
	fruits[3] = "grape"
	fruits[4] = "plum"

	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	nums := [5]int{1, 2, 3, 4, 5}

	for i, num := range nums {
		fmt.Println(i, num)
	}
}

func RangeDemo() {
	fmt.Println("-- Using index range --")
	friends := [5]string{"Annie", "Bethesda", "Charlie", "Dylan", "Edmund"}
	fmt.Printf("Before: %s\n", friends[1])

	for i := range friends {
		friends[1] = "Bilbo"

		if i == 1 {
			fmt.Printf("After: %s\n", friends[1])
		}
	}

	fmt.Println("-- Using index and value range --")
	friends = [5]string{"Annie", "Bethesda", "Charlie", "Dylan", "Edmund"}
	fmt.Printf("Before: %s\n", friends[1])
	for i, v := range friends {
		friends[1] = "Butter"

		if i == 1 {
			// we are iterating over a copy of friends!
			// Even though the value at 1 has changed,
			// v is still the original value ('Bethesda')
			fmt.Printf("After (index): %s\n", friends[1])
			fmt.Printf("After (value): %s\n", v)
		}
	}
}
