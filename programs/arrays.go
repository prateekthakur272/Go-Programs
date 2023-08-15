package programs

import "fmt"

func Arrays() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	var integers = [...]int{1, 2, 3}
	names := [3]string{"Prateek", "Tanu", "Palak"}
	fmt.Println(numbers)
	fmt.Println(integers)
	fmt.Println(names[0])

	numbers[0] = 50
	fmt.Println(numbers)

	var positions = [5]int{0: 4, 3: 6}
	fmt.Println(positions)
	fmt.Println(len(positions))
	fmt.Println(len(integers))
}
