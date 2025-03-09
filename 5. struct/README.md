# Go Structs Tutorial

## Overview

Structs in Go are composite data types that group together variables under a single name. They are used to create more complex data structures.

## Types of Structs

### 1. Basic Struct

A basic struct is a collection of fields.

```go
type Person struct {
    name string
    age  int
}

// Method for Person
func (p Person) NameOfPerson() string {
    return p.name
}
```

### 2. Nested Struct

A nested struct is a struct that contains another struct as a field.

```go
type Address struct {
    street  string
    city    string
    country string
}

type Employee struct {
    Person   Person // regular nesting
    Address         // promoted/embedded struct
    salary   int
    contacts []string
}
```

### 3. Struct with Function Field

A struct can have a function as a field.

```go
type Calculator struct {
    nums    []int
    operate func([]int) int
}
```

### 4. Struct with Interface Field

A struct can have an interface as a field, allowing for polymorphism (ability of different types).

```go
type Printer interface {
    Print()
    Change()
}

type Document struct {
    text string
}

func (d Document) Change(s string) {
    d.text = s
}

func (d Document) Print() {
    fmt.Println(d.text)
}
```

### 5. Anonymous Field Struct

A struct can have anonymous fields, where the type of the field is used as the field name.

```go
type Student struct {
    string // anonymous field
    int    // anonymous field
    grades []int
}
```

## Initialization Methods

### 1. Named Fields

```go
person := Person{
    name: "Abul",
    age:  10,
}
```

### 2. Ordered Fields

```go
person := Person{"Babul", 20} // order matters
```

### 3. Zero Value

```go
var person Person // all fields will have zero values
```

### 4. Pointer to Struct

```go
person := &Person{name: "Cabul", age: 30}
```

### 5. Anonymous Struct

```go
point := struct {
    X, Y int
}{
    X: 10,
    Y: 20,
}
```

## Access and Modification

### 1. Direct Access

```go
person.name = "Abul"
fmt.Println(person.age)
```

### 2. Pointer Access

```go
personPtr := &person
(*personPtr).name = "Babul"  // explicit
personPtr.name = "Babul"     // implicit
```

### 3. Embedded Fields

```go
student := Student{
    Person: Person{name: "Dabul", age: 40},
    Grade:  95,
}
fmt.Println(student.name)  // directly access Person's name
fmt.Println(student.Person.name)  // explicit access
```

## Common Operations

### 1. Comparing Structs

- Structs are comparable if all fields are comparable.
- Use `==` for equality comparison.
- Deep comparison may need custom implementation.

```go
p1 := Person{"Abul", 10}
p2 := Person{"Abul", 10}
fmt.Println(p1 == p2) // true
```

### 2. Copying Structs

- Structs are copied by value.
- Pointer fields share the same memory.

```go
p1 := Person{"Abul", 10}
p2 := p1
p2.name = "Babul"
fmt.Println(p1.name) // Abul
fmt.Println(p2.name) // Babul
```

### 3. String Representation

```go
fmt.Printf("%v\n", person)    // {Abul 10}
fmt.Printf("%+v\n", person)   // {name:Abul age:10}
fmt.Printf("%#v\n", person)   // main.Person{name:"Abul", age:10}
```

## Memory Considerations

| Type      | When to Use                  |
| --------- | ---------------------------- |
| Value     | Small structs, need copies   |
| Pointer   | Large structs, need mutation |
| Embedded  | Code reuse, composition      |
| Anonymous | One-time use                 |
