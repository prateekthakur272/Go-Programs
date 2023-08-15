package programs

import "fmt"

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
}
