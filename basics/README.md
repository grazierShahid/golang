# Go Basics Tutorial

## 1. Package and Import

```go
package main
import "fmt"
import (
    "strings"
    "math"
)
```

## Keywords

Go has a set of reserved keywords that cannot be used as identifiers. Some of the keywords are:

- `break`, `default`, `func`, `interface`, `select`
- `case`, `defer`, `go`, `map`, `struct`
- `chan`, `else`, `goto`, `package`, `switch`
- `const`, `fallthrough`, `if`, `range`, `type`
- `continue`, `for`, `import`, `return`, `var`

## 2. Data Types

| Category | Types                                                                     | Zero Value |
| -------- | ------------------------------------------------------------------------- | ---------- |
| Numeric  | `int`, `int8`...`int64`, `uint`, `uint8`...`uint64`, `float32`, `float64` | `0`        |
| Text     | `string`, `rune`, `byte`                                                  | `""`       |
| Boolean  | `bool`                                                                    | `false`    |
| Complex  | `complex64`, `complex128`                                                 | `0+0i`     |
| Pointer  | `*T`                                                                      | `nil`      |

## Variables

Variables in Go can be declared using the `var` keyword or the shorthand `:=` operator.

```go
var a int = 10
var b = 20
c := 30
```

## 3. Variables

```go
var name string = "John"    // explicit type
var age = 25               // type inference
score := 95               // shorthand

// Multiple declarations
var a, b, c = 1, 2, 3
var (
    name   string
    age    int
    isValid bool
)
```

## Constants

Constants are declared using the `const` keyword and cannot be changed after they are set.

```go
const Pi = 3.14
const Greeting = "Hello, World!"
```

## 4. Constants

```go
const Pi = 3.14
const (
    StatusOK = 200
    StatusNotFound = 404
)

// iota for incremental constants
const (
    Sunday = iota    // 0
    Monday          // 1
    Tuesday         // 2
)
```

## 5. Basic Operations

| Category   | Operators                         | Example  |
| ---------- | --------------------------------- | -------- |
| Arithmetic | `+`, `-`, `*`, `/`, `%`           | `a + b`  |
| Relational | `==`, `!=`, `<`, `<=`, `>`, `>=`  | `a == b` |
| Logical    | `&&`, `\|\|`, `!`                 | `a && b` |
| Bitwise    | `&`, `\|`, `^`, `&^`, `<<`, `>>`  | `a & b`  |
| Assignment | `=`, `+=`, `-=`, `*=`, `/=`, `%=` | `a += b` |

## Scope

Scope in Go is determined by where variables and constants are declared. There are three main scopes:

- **Package scope**: Variables declared outside functions are visible throughout the package.
- **Function scope**: Variables declared inside a function are visible only within that function.
- **Block scope**: Variables declared inside a block (e.g., `if`, `for`, `switch`) are visible only within that block.

## 6. Type Casting

```go
var i int = 42
var f float64 = float64(i)
var s string = string(i)
```

## 7. Scope

| Level    | Visibility           | Example                      |
| -------- | -------------------- | ---------------------------- |
| Package  | All files in package | `var PackageVar = 100`       |
| Function | Inside function      | `func main() { var x = 10 }` |
| Block    | Inside block `{}`    | `if true { var y = 20 }`     |

## 8. Comments

```go
// Single-line comment

/*
   Multi-line
   comment
*/
```

## 9. Print Functions

```go
fmt.Print("No newline")
fmt.Println("With newline")
fmt.Printf("Formatted: %d, %s", 42, "text")
```

## Running the Code

To run the code, use the following command:

```bash
go run main.go
```
