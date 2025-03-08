# Go Functions and Methods Tutorial

## Functions

### 1. Basic Functions

Functions are the basic building blocks of Go programs. They can take parameters and return values.

```go
func add(a, b int) int {
    return a + b
}
```

### 2. Variadic Functions

Variadic functions can accept a zero to multiple parameter of same type.

```go
func sum(nums ...int) int {}
```

### 3. Anonymous Functions

Anonymous functions are created for one-time use or one place only. They can be written in different ways:

```go
// Inside a function
func() {}()

// Assigning to a variable
f := func() {}

// Taking anonymous function as argument
func executeFunc(fn func(a, b string) string) {
    fmt.Println(fn("hello ", "world"))
}

// Returning anonymous function
func f2() func(a, b int) int {
    return func(a, b int) int {
        return a + b
    }
}
```

### 4. Pass By Value and Reference

Go supports both pass by value and pass by reference.

```go
func passBy(a int, b *int) {}
```

### 5. Multiple Return Values

Functions in Go can return multiple values.

```go
func f3(a, b int) (sum, mul int) {}
```

### 6. Higher-Order Functions

Higher-order functions are functions that take other functions as parameters or/and return functions.

```go
func higherOrder(fn func(int, int) int) func(int) int {
    return func(x int) int {
        return fn(x, x)
    }
}
```

### 7. Init Function

The `init` function always runs first if it exists. If there are multiple `init` functions, they execute in the order of placement. The `init` function cannot have any parameters or return values.

```go
func init() {
    fmt.Println("let me execute first, side please....")
}
```

## Recursive Functions

A recursive function is one that calls itself to solve a smaller instance of the same problem. Every recursive function needs:

1. Base case(s) to stop recursion
2. Recursive case that moves toward the base case

```go
// Example 1: Factorial - O(n)
func factorial(n int) int {
    if n <= 1 {          // Base case
        return 1
    }
    return n * factorial(n-1)  // Recursive case
}

// Example 2: Fibonacci - O(2^n)
func fibonacci(n int) int {
    if n <= 1 {          // Base cases
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)  // Recursive case
}

// Example 3: Array Sum - O(n)
func sumRecursive(nums []int) int {
    if len(nums) == 0 {  // Base case
        return 0
    }
    return nums[0] + sumRecursive(nums[1:])  // Recursive case
}
```

### Time Complexity Analysis

| Function  | Complexity | Reason                                        |
| --------- | ---------- | --------------------------------------------- |
| Factorial | O(n)       | Makes n recursive calls                       |
| Fibonacci | O(2^n)     | Makes two recursive calls for each n          |
| Array Sum | O(n)       | Makes n recursive calls, one for each element |

## Methods

Methods are functions associated with a specific type. The difference between a function and a method is that in a method, you need to mention the receiver (which type it belongs to).

### Methods are specified in some types:

1. **Struct Type**

```go
type User struct {
    Name string
    Age  int
}

func (u User) who() {
    fmt.Println("this is", u.Name, "and he is", u.Age, "years old.")
}
```

2. **Non-struct Type**

```go
type num int

func (n num) square() num {
    return n * n
}
```

3. **Pointer Receiver**

```go
func (u *User) changeName() {
    u.Name = "Shahid"
}
```

4. **Accepting Both Pointer and Value**

```go
func (u *User) changeNameTo(name string) {
    u.Name = name
}
```

### Important Notes:

- You can write multiple methods with the same name but for different types of receivers.
- You cannot write multiple functions with the same name.
- Example:

  ```go
  func (u User) whatisthenum() int {
      return u.Age
  }

  func (n num) whatisthenum() int {
      return n
  }
  ```

## Differences Between Functions and Methods

| Aspect               | Functions                                     | Methods                                                      |
| -------------------- | --------------------------------------------- | ------------------------------------------------------------ |
| Definition           | Standalone functions                          | Functions associated with a type                             |
| Receiver             | No receiver                                   | Requires a receiver                                          |
| Syntax               | `func name(params) returnType`                | `func (receiver Type) name(params) returnType`               |
| Usage                | Called directly                               | Called on an instance of the type                            |
| Multiple Definitions | Cannot have multiple functions with same name | Can have multiple methods with same name for different types |

## Running the Code

To run the code, use the following command:

```bash
go run main.go
```
