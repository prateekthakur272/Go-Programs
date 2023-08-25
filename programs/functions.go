package programs

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloTo(name string) {
	fmt.Println("Hello ", name)
}

func sum(x int, y int) int {
	return x + y
}

func diff(x int, y int) (result int) {
	result = x - y
	return
}

func count(x int, n int) int {
	fmt.Println(x)
	if x < n {
		return count(x+1, n)
	}
	return 0
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func Functions() {
	sayHello()
	sayHelloTo("Prateek")
	count(1, 10)
	fmt.Println(factorial(5))
}
