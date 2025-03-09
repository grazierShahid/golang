package main

import "fmt"

// 1. Basic struct
type Person struct {
	name string
	age  int
}

// Method for Person
func (p Person) NameOfPerson() string {
	return p.name
}

// 2. Nested struct
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

// 3. Struct with function field
type Calculator struct {
	nums    []int
	operate func([]int) int
}

// 4. Struct with interface field
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

// 5. Anonymous struct field
type Student struct {
	string // anonymous field
	int    // anonymous field
	grades []int
}

func main() {
	// 1. Basic struct usage
	p1 := Person{"Abul", 20}
	p2 := Person{name: "Babul"} // partial initialization

	fmt.Printf("Person-1: %+v\n", p1)
	fmt.Printf("Person-2: %+v\n", p2)
	fmt.Println("Name of Person-1:", p1.NameOfPerson())

	// 2. Struct equality
	p3 := Person{"Abul", 20}
	fmt.Printf("is p1 == p3 ? : %v\n", p1 == p3)
	fmt.Printf("is p1 == p2 ? : %v\n", p1 == p2)

	// 3. Nested struct
	emp := Employee{
		Person: Person{"Cabul", 27},
		Address: Address{
			street:  "Begum Rokeya Sarani",
			city:    "Dhaka",
			country: "Bangladesh",
		},
		salary:   50000,
		contacts: []string{"+8801xxxxxx", "cabul@email.com"},
	}

	// Accessing nested fields
	fmt.Printf("Employee name: %s\n", emp.Person.name) // can't access as emp.name
	fmt.Printf("Employee city: %s\n", emp.city)        // direct access

	// 4. Anonymous struct
	length := struct {
		l, h int
	}{ // can write as {10, 20} too
		h: 10,
		l: 20,
	}
	fmt.Printf("Point: %+v\n", length)

	// 5. Struct with function field
	calc := Calculator{
		nums: []int{1, 2, 3, 4, 5},
		operate: func(nums []int) int {
			sum := 0
			for _, n := range nums {
				sum += n
			}
			return sum
		},
	}
	a := []int{1, 2, 3, 4}
	fmt.Printf("Sum: %v\n", calc.operate(a))
	fmt.Printf("Sum: %v\n", calc.operate(calc.nums))

	// 6. Anonymous field usage
	student := Student{
		string: "Dabul",
		int:    20,
		grades: []int{40, 60, 80, 90},
	}
	fmt.Printf("Student: %+v\n", student)

	// 7. Pointer to struct
	empPtr := &emp
	fmt.Printf("Employee salary: %d\n", empPtr.salary)

	// 8. Struct slice
	people := []Person{
		{"Ebul", 10},
		{"Fbul", 20},
		{"Gbul", 30},
	}
	fmt.Printf("People: %+v\n", people)

	// 9. Struct map
	employeeMap := map[string]Employee{
		"IT": emp,
	}
	fmt.Printf("IT Employee: %+v\n", employeeMap["IT"])
}
