package main

import "fmt"

func main() {
	// 1. Variable declarations
	var a int = 10 // explicit type
	var b = 20     // type inference
	c := 30        // shorthand

	// Multiple declarations
	var d, e, f = 1, 2, 3
	var (
		name    string = "John"
		age     int    = 25
		isValid bool   = true
	)

	// 2. Constants
	const Pi = 3.14
	const (
		StatusOK       = 200
		StatusNotFound = 404
	)

	// Constants with iota
	const (
		Sunday = iota
		Monday
		Tuesday
	)

	// 3. Basic types
	var (
		integer int       = 42
		float   float64   = 3.14
		str     string    = "hello"
		boolean bool      = true
		complex complex64 = 1 + 2i
	)

	// 4. Type casting
	var x int = 65
	var y float64 = float64(x)
	var z string = string(x) // ASCII value 'A'

	// 5. Basic operations
	fmt.Println("--- Operations ---")
	fmt.Println("Arithmetic:", a+b, a-b, a*b, a/b, a%b)
	fmt.Println("Compare:", a == b, a != b, a < b)
	fmt.Println("Logical:", true && false, true || false, !true)
	fmt.Println("Bitwise:", a&b, a|b, a^b, a<<1, a>>1)

	// 6. Print formatting
	fmt.Printf("Types: %T %T %T\n", x, y, z)
	fmt.Printf("Values: %d %.2f %s\n", x, y, z)
}
