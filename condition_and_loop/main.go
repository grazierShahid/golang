package main

import "fmt"

func main() {
	// 1. If-Else Conditions
	x := 10

	// Basic if
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// If with initialization
	if y := 20; y > x {
		fmt.Println("y is greater than x")
	}

	// If-else
	if x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}

	// If-else if-else
	score := 85
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

	// 2. Switch Statements
	day := "Monday"

	// Basic switch
	switch day {
	case "Monday":
		fmt.Println("Start of week")
	case "Friday":
		fmt.Println("End of week")
	default:
		fmt.Println("Regular day")
	}

	// Switch with multiple cases
	n := 2
	switch n {
	case 1, 3, 5:
		fmt.Println("Odd")
	case 2, 4, 6:
		fmt.Println("Even")
	}

	// Switch without expression
	switch {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}

	// 3. Loops
	// Basic for loop
	for i := 0; i < 3; i++ {
		fmt.Println("Count:", i)
	}

	// While-like loop
	j := 0
	for j < 3 {
		fmt.Println("While-like:", j)
		j++
	}

	// Infinite loop with break
	count := 0
	for {
		if count >= 3 {
			break
		}
		fmt.Println("Infinite:", count)
		count++
	}

	// Continue statement
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("Skipping 2:", i)
	}

	// Range loop with slice
	numbers := []int{1, 2, 3}
	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	// Range with map
	colors := map[string]string{"red": "#ff0000", "green": "#00ff00"}
	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Range with string
	for i, char := range "Hello" {
		fmt.Printf("Position: %d, Character: %c\n", i, char)
	}
}
