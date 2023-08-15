package programs

import (
	"fmt"
)

func Operators() {
	var (
		a int     = 5
		b int     = 10
		c bool    = true
		d float32 = 3.14
	)
	fmt.Printf("a:%v\nb:%v\nc:%v\nd:%v\n", a, b, c, d)
	// Arithmatic
	fmt.Printf("a+b=%v\n", a+b)
	fmt.Printf("b-a=%v\n", b-a)
	fmt.Printf("a*b=%v\n", a*b)
	fmt.Printf("b/a=%v\n", b/a)
	fmt.Printf("b%%a=%v\n", b%a)
	a++
	b--
	fmt.Printf("a:%v\nb:%v\nc:%v\nd:%v\n", a, b, c, d)

	// Assignment
	var x = 10
	x += 5
	fmt.Printf("x+=5 :%v\n", x)
	x = 10
	x -= 6
	fmt.Printf("x-=6 :%v\n", x)
	x = 10
	x *= 4
	fmt.Printf("x*=4 :%v\n", x)
	x = 10
	x /= 2
	fmt.Printf("x/=2 :%v\n", x)
	x = 10
	x %= 3
	fmt.Printf("x%%=3 :%v\n", x)
	x = 10
	x &= 3
	fmt.Printf("x&=3 :%v\n", x)
	x = 10
	x |= 3
	fmt.Printf("x|=3 :%v\n", x)
	x = 10
	x ^= 3
	fmt.Printf("x^=3 :%v\n", x)
	x = 10
	x >>= 2
	fmt.Printf("x>>=2 :%v\n", x)
	x = 10
	x <<= 2
	fmt.Printf("x<<=2 :%v\n", x)

	var (
		p int = 5
		q int = 8
	)
	// Comparision
	fmt.Printf("p>6: %t\n", p > 4)
	fmt.Printf("q<10: %t\n", q < 10)
	fmt.Printf("p==5: %t\n", p == 5)
	fmt.Printf("p>=5: %t\n", p >= 5)
	fmt.Printf("q<=4: %t\n", p <= 4)

	// Logical
	var (
		t = true
		f = false
	)
	fmt.Printf("true && true: %t\n", true && t)
	fmt.Printf("true && false: %t\n", t && f)
	fmt.Printf("true || true: %t\n", true || t)
	fmt.Printf("true || false: %t\n", t || f)
	fmt.Printf("!true: %t\n", !t)

	// Bitwise
	x = 6
	fmt.Printf("%b\n", x)
	fmt.Printf("x & 2 = %03b\n", x&2)
	fmt.Printf("x | 2 = %03b\n", x|2)
	fmt.Printf("x ^ 2 = %03b\n", x^2)
	fmt.Printf("x << 2 = %03b\n", x<<2)
	fmt.Printf("x >> 2 = %03b\n", x>>2)
}
