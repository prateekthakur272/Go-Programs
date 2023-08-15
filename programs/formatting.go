package programs

import "fmt"

func Formatting() {

	var message string = "hello"
	var a = 13
	var b = 2.126428
	var c = -2

	fmt.Println(message, a, b)

	// General Formatting Verbs
	// Prints the value in the default format
	fmt.Printf("%v\n", message)
	// Prints the value in Go-syntax format
	fmt.Printf("%T\n", message)
	// Prints the type of the value
	fmt.Printf("%#v\n", message)
	// Prints the % sign
	fmt.Printf("%%\n")

	// Integer Formatting Verbs
	// Base 2
	fmt.Printf("%b\n", a)
	// Base 10
	fmt.Printf("%d\n", a)
	// Base 10 and always show sign
	fmt.Printf("%+d\n", c)
	// Base 8
	fmt.Printf("%o\n", a)
	// Base 8, with leading 0o
	fmt.Printf("%O\n", a)
	// Base 16, lowercase
	fmt.Printf("%x\n", a)
	// Base 16, uppercase
	fmt.Printf("%X\n", a)
	// Base 16, with leading 0x
	fmt.Printf("%#X\n", a)
	// Pad with spaces (width 4, right justified)
	fmt.Printf("%4d\n", a)
	// Pad with spaces (width 4, left justified)
	fmt.Printf("%-4d\n", a)
	// Pad with zeroes
	fmt.Printf("%04d\n", a)

	// String Formatting Verbs
	// Prints the value as plain string
	fmt.Printf("%s\n", message)
	// Prints the value as a double-quoted string
	fmt.Printf("%q\n", message)
	// Prints the value as plain string (width 8, right justified)
	fmt.Printf("%8s\n", message)
	// Prints the value as plain string (width 8, left justified)
	fmt.Printf("%-8s\n", message)
	// Prints the value as plain string (width 8, left justified)
	fmt.Printf("%x\n", message)
	// Prints the value as plain string (width 8, left justified)
	fmt.Printf("% x\n", message)

	// Boolean Formatting Verbs
	// Value of the boolean operator in true or false format (same as using %v)
	fmt.Printf("%t\n", true)

	// Float Formatting Verbs
	// Scientific notation with 'e' as exponent
	fmt.Printf("%e\n", b)
	// Decimal point, no exponent
	fmt.Printf("%f\n", b)
	// Default width, precision 2
	fmt.Printf("%.2f\n", b)
	// Width 6, precision 2
	fmt.Printf("%6.2f\n", b)
	// Exponent as needed, only necessary digits
	fmt.Printf("%g\n", b)

}
