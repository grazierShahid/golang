package main

import "fmt"

// Basic function
func add(a, b int) int { // same as add(a int, b int)
	return a + b
}

// Variadic function
func sum(nums ...int) int { // ...int is for any num of value, like nums is an array
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Anonymous function as argument
func executeFunc(fn func(a, b string) string) {
	fmt.Println(fn("pera ", "nai, "))
}

// Function returning anonymous function
func f2() func(a, b int) int {
	return func(a, b int) int {
		return a + b
	}
}

// Pass by value and reference
func passBy(a int, b *int) { // a - value; b - reference
	a = a * 2
	*b = *b * 2
}

// Multiple return values
func f3(a, b int) (sum, mul int) { // sum and mul initialized here
	sum = a + b // dont need to initialize as 'var sum' or using (:=)
	mul = a * b
	return // you may set write as return sum, mul. but already written in return list
}

// Struct type
type User struct {
	Name string
	Age  int
}

// Method with value receiver
func (u User) who() {
	fmt.Println("this is", u.Name, "and he is", u.Age, "years old.")
}

// Non-struct type
type num int

// Method on non-struct type
func (n num) square() num {
	return n * n
}

// Pointer receiver method
func (u *User) changeName() {
	u.Name = "Shahid"
}

// Method accepting both pointer and value
func (u *User) changeNameTo(name string) {
	u.Name = name
}

// Higher-order function = recive a func or/and return a func
func higherOrder(fn func(int, int) int) func(int) int {
	return func(x int) int {
		return fn(x, x)
	}
}

// Init function
func init() {
	fmt.Println("let me execute first, side please....")
}

// Recursive function examples
func factorial(n int) int {
	// Base case
	if n <= 1 {
		return 1
	}
	// Recursive case
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	// Base cases
	if n <= 1 {
		return n
	}
	// Recursive case
	return fibonacci(n-1) + fibonacci(n-2)
}

func sumRecursive(nums []int) int {
	// Base case
	if len(nums) == 0 {
		return 0
	}
	// Recursive case
	return nums[0] + sumRecursive(nums[1:])
}

func main() {
	// Calling variadic function
	fmt.Println(sum(1, 2, 3, 4))

	// Passing anonymous function as argument
	anoy := func(a, b string) string {
		return a + b + "chill"
	}
	executeFunc(anoy)

	// Taking an anonymous function as return
	value := f2()
	fmt.Println(value(2, 3))

	aa, bb := 2, 3
	defer passBy(aa, &bb) // defer function, bb will updated (referece), aa will not (value)
	fmt.Println(aa, bb)   // now: 2, 3 (bcoz passBy() called in defer)

	s, m := f3(aa, bb)
	fmt.Println(s, m)

	person := User{
		Name: "grazier",
		Age:  72,
	}
	person.who()

	var n num = 5
	fmt.Println(n.square())

	person.changeName()
	person.who()

	person.changeNameTo("Hasan")
	person.who()

	// Higher-order function usage
	double := higherOrder(add)
	fmt.Println(double(5))

	// Recursive function examples
	fmt.Println("\n--- Recursive Functions ---")
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Printf("Fibonacci of 6: %d\n", fibonacci(6))
	fmt.Printf("Sum of [1,2,3,4,5]: %d\n", sumRecursive([]int{1, 2, 3, 4, 5}))
}
