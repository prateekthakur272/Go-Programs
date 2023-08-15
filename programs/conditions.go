package programs

import "fmt"

func Conditions() {
	var (
		a int = 10
		b int = 10
	)
	// Conditions
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a != b)
	// If Statement
	if a > b {
		fmt.Println("a is greater than b")
	}
	// If-Else Statement
	if a < b {
		fmt.Println("a is smaller than b")
	} else {
		fmt.Println("a is greater than or equal to b")
	}

	// If-Else-If Statement
	if a > b {
		fmt.Println("a greater than b")
	} else if a < b {
		fmt.Println("a is smaller than b")
	} else {
		fmt.Println("a is equal to b")
	}

	// Nested If Statement
	if a > 4 {
		fmt.Println("a is greater than 4")
		if a > 8 {
			fmt.Println("a is greater than 8 also")
		}
	}
}
