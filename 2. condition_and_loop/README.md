# Go Conditions and Loops Tutorial

## Conditions

### 1. If Statement

```go
// Basic if
if condition {
    // code
}

// If with initialization
if x := value; condition {
    // code
}

// If-else
if condition {
    // code
} else {
    // code
}

// If-else if-else
if condition1 {
    // code
} else if condition2 {
    // code
} else {
    // code
}
```

### 2. Switch Statement

```go
// Basic switch
switch variable {
case value1:
    // code
case value2:
    // code
default:
    // code
}

// Multiple cases
switch variable {
case 1, 2, 3:
    // code
default:
    // code
}

// Switch without expression
switch {
case condition1:
    // code
case condition2:
    // code
}
```

## Loops

### 1. For Loop Types

| Type       | Syntax                                 | Use Case                        |
| ---------- | -------------------------------------- | ------------------------------- |
| Basic      | `for i := 0; i < n; i++`               | Counter-based iteration         |
| While-like | `for condition {}`                     | Condition-based iteration       |
| Infinite   | `for {}`                               | Continuous execution with break |
| Range      | `for index, value := range collection` | Iterating collections           |

### 2. Control Statements

- `break`: Exit the loop
- `continue`: Skip to next iteration
- `goto`: Jump to labeled statement (use sparingly)

### 3. Range Usage

```go
// Slice/Array
for index, value := range slice {}

// Map
for key, value := range map {}

// String
for index, char := range string {}

// Channel
for value := range channel {}
```
