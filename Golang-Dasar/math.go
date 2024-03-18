package main

import "fmt"

func main() {
	// Operator
	var (
		a = 10
		b = 20
		c = a + b
		d = a - b
		e = a * b
		f = a / b
		g = a % b
	)

	// Augmented Assignments
	// +=, -=, *=, /=. %=
	var (
		i = 10
	)
	i += 10

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(i)

	i += 5
	fmt.Println(i)

	// Unary Operator
	// ++, --. -, +, !
	var (
		v = 10
	)
	v++
	fmt.Println(v)

	v--
	fmt.Println(v)

}
