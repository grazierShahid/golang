# Go Arrays and Slices Tutorial

## Introduction

Arrays and slices are two closely related data structures in Go:

- **Arrays**: Fixed-size collections of elements with the same type
- **Slices**: Dynamic, flexible views into arrays with automatic resizing capabilities

## Table of Contents

1. [Arrays](#arrays)
2. [Slices](#slices)
3. [Higher-Order Functions](#higher-order-functions)
4. [Code Examples](#code-examples)

## Arrays

- Array declaration and initialization
- Multi-dimensional arrays
- Value semantics vs. pointer semantics
- Array comparison
- Iterating through arrays with range

## Slices

- Slice creation and initialization
- Relationship between length and capacity
- Dynamic resizing with append
- Slicing operations
- Sorting and searching
- Working with byte slices
- Multi-dimensional slices
- Slices of structs

## Higher-Order Functions

- Transformation functions
- Filtering operations
- Custom comparisons

## Code Examples

### Arrays vs. Slices

```go
// Arrays have fixed size
var arr [5]int

// Slices have dynamic size
var sli []int
```

### Slice Length and Capacity

Key concepts:

- **Length**: Number of elements in the slice
- **Capacity**: Maximum number of elements before reallocation

When capacity is exceeded, Go automatically doubles the capacity, similar to C++ vectors.

### Value vs. Reference Semantics

```go
// Arrays are copied on assignment
source := [5]int{10, 20, 30, 40, 50}
dest := source  // Creates a copy

// Slices share the same underlying array
sourceSlice := []int{10, 20, 30, 40, 50}
destSlice := sourceSlice  // Points to the same underlying array
```

### Common Slice Operations

- Appending elements
- Taking sub-slices
- Sorting
- Searching
- Filtering

### Higher-Order Function Examples

The tutorial demonstrates using functions as arguments for:

- Mapping operations (`makeSquare`)
- Filtering data (`numFilter`)
- Finding elements (`whoIsYounger`)

## Code Structure

The `main.go` file demonstrates concepts in this order:

1. Basic array operations
2. Multi-dimensional arrays
3. Array semantics
4. Array comparisons
5. Basic slice operations
6. Slice length and capacity
7. Slice manipulation
8. Advanced slice operations
9. Higher-order functions
