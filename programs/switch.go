package programs

import (
	"fmt"
)

func Switch() {
	// Single Case Switch Statement
	var x int = 7
	switch x {
	case 1:
		fmt.Println("JANUARY")
	case 2:
		fmt.Println("FEBRUARY")
	case 3:
		fmt.Println("MARCH")
	case 4:
		fmt.Println("APRIL")
	case 5:
		fmt.Println("MAY")
	case 6:
		fmt.Println("JUNE")
	case 7:
		fmt.Println("JULY")
	case 8:
		fmt.Println("AUGUST")
	case 9:
		fmt.Println("SEPTEMBER")
	case 10:
		fmt.Println("OCTOBER")
	case 11:
		fmt.Println("NOVEMBER")
	case 12:
		fmt.Println("DECEMBER")
	default:
		fmt.Println("Invalid Input")
	}

	// Multi Case Switch Statement

	switch x {
	case 2:
		fmt.Println("28 Days (29 Days in leap year)")
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 Days")
	case 4, 6, 9, 11:
		fmt.Println("30 Days")
	}
}
