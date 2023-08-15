package programs

import "fmt"

func Print() {
	var (
		message1 string = "Hello"
		message2 string = "GoLang"
	)

	print(message1)
	print(message2, '\n')

	println(message1)
	println(message2)

	fmt.Printf("message1: %v %T\n", message1, message1)
	fmt.Printf("message2: %v %T\n", message2, message2)

	fmt.Printf("message1: %#v", message1)
}
