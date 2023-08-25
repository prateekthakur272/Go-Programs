package programs

import (
	"fmt"
)

func Loops() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		} else if i == 7 {
			break
		}
		fmt.Println(i)
	}

	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			fmt.Println(a, b)
		}
	}

	var fruits = [5]string{"Apple", "Banana", "Orange", "Kiwi", "Pineapple"}
	for index, name := range fruits {
		fmt.Println(index, name)
	}
}
