package programs

import (
	"fmt"
)

func Slices() {
	// Slices
	integers := []int{}
	fmt.Println(integers)
	integers = append(integers, 1)
	fmt.Println(integers)
	fmt.Println(cap(integers), len(integers))

	// Slice from array
	numbers := [...]int{3, 4, 5, 6, 7}
	fmt.Println(numbers)
	num_slice := numbers[0:]
	fmt.Println(cap(numbers), len(numbers))
	fmt.Println(num_slice)

	// Slice using make function
	positions := make([]int, 5, 5)
	fmt.Println(positions, len(positions), cap(positions))
	rank := make([]int, 5)
	fmt.Println(rank, len(rank), cap(rank))

	// Modify
	positions[0] = 10
	positions[1] = 20
	positions[2] = 30
	fmt.Println(positions[2])
	fmt.Println(positions)

	// Append
	positions = append(positions, 7)
	fmt.Println(positions)
	fmt.Println(append(positions, num_slice...))

	array := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice := array[4:7]
	fmt.Println(slice, len(slice), cap(slice))

	// Copy
	large := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(large, len(large), cap(large))
	small := large[5:10]
	fmt.Println(small, len(small), cap(small))
	small_array := make([]int, len(small))
	fmt.Println(small_array, len(small_array), cap(small_array))
	copy(small_array, small)
	fmt.Println(small_array, len(small_array), cap(small_array))
}
