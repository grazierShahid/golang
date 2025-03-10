package main

import (
	"bytes"
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Array Declaration
	var arr1 [5]int // [0, 0, 0, 0, 0]
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{1, 2, 3}
	arr4 := [5]int{1: 5, 3: 10}

	fmt.Println(arr1, arr2, arr3, arr4) // [0 0 0 0 0] [1 2 3 4 5] [1 2 3] [0 5 0 10 0]

	// Multi-dimensional array
	arr2d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2d) // [[1 2 3] [4 5 6]]
	for i, row := range arr2d {
		for j, x := range row {
			fmt.Print(x) // 123456
			arr2d[i][j] *= 2
		}
	}
	fmt.Println(arr2d) //[[2 4 6] [8 10 12]]

	// Array value vs pointer
	source := [5]int{10, 20, 30, 40, 50}
	dest1 := source  // value copy
	dest2 := &source // pointer reference / address copy

	dest1[0] = 111
	dest2[1] = 222                    // same as (*dest2)[1] = 222
	fmt.Println(source, dest1, dest2) // [10 222 30 40 50] [111 20 30 40 50] &[10 222 30 40 50]

	// Array comparison
	arr5 := [3]int{1, 2, 3}
	arr6 := [3]int{1, 2, 3}
	arr7 := [3]int{1, 2, 5}
	fmt.Println(arr5 == arr6, arr6 == arr7) // true false; must be the same length of both array besid ==

	/// Slices

	// array is fixed sized; slice is dynamic
	var sli1 []int
	sli2 := []int{1, 2, 3}
	sli3 := make([]int, 3, 5)     // len = 3, cap = 5
	fmt.Println(sli1, sli2, sli3) // [] [1 2 3] [0 0 0]

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	subset := numbers[1:5]
	sli2 = append(sli2, 44, 55)
	fmt.Println(subset, sli2) // [2 3 4 5] [1 2 3 44 55]

	sli3[0] = 11
	sli3 = append(sli3, 11, 22)             // where those will append? len(sli3) = ? and cap(sli3) = ?
	sli3 = append(sli3, 33)                 // now len(sli3) = ? and cap(sli3) = ? lets see...
	fmt.Println(sli3, len(sli3), cap(sli3)) // [11 0 0 11 22 33] 6 10
	// capacity will increase by double after crossing the len(); just like vector in C++

	// lets see one more time about the len and cap
	s := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Println(s, len(s), cap(s)) // [0 1 2 3 4 5 6 7 8 9] 10 10
	s = append(s, 11)
	fmt.Println(s, len(s), cap(s)) // [0 1 2 3 4 5 6 7 8 9 11] 11 20

	// lets DEEP-DIVE in advance
	// sorting
	nums := []int{2, 4, 4, 1, 1, 3, 9}
	sort.Ints(nums)   // standard sort
	fmt.Println(nums) // [1 1 2 3 4 4 9]

	sort.Slice(nums, func(i, j int) bool { // custom anoymous func
		return nums[i] > nums[j] // descending
	})
	fmt.Println(nums) // [9 4 4 3 2 1 1]

	// searching
	sort.Ints(nums) // must be sorted in ascending order
	ind := sort.SearchInts(nums, 3)
	fmt.Println(ind) // 3 (0-based indexing)

	// Byte slice operations
	text := []byte("this is grazier!")
	upper := bytes.ToUpper(text)
	trimmed := bytes.TrimSpace(text)
	split := bytes.Split(text, []byte(" "))
	fmt.Println(text, upper, trimmed, split)                  // output will be ASCII code for each char
	fmt.Println(string(text), string(upper), string(trimmed)) // this is grazier! THIS IS GRAZIER! this is grazier!
	// split cannot print as string(split). whatever! to see its value as char
	for _, char := range split {
		fmt.Println(string(char)) // this \n is \n grazier!
	}

	// slice of struct
	people := []Person{
		{"Abul", 20},
		{"Babul", 50},
		{"Cabul", 10},
		{"Dabul", 40},
		{"Ebul", 30},
	}

	fmt.Println(people) // [{Abul 20} {Babul 50} {Cabul 10} {Dabul 40} {Ebul 30}]
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age // ascending by age wise
	})
	fmt.Println(people) // [{Cabul 10} {Abul 20} {Ebul 30} {Dabul 40} {Babul 50}]

	// multi-dimensional slice
	mat := make([][]int, 2) // row : 2
	for i := range mat {
		mat[i] = make([]int, 3) // colm : 3
		for j := range mat[i] {
			mat[i][j] = i + j
		}
	}
	fmt.Println(mat) // [[0 1 2] [1 2 3]]

	// array vs slice relation

	// when we copy a slice from an array, then actually copied its memory address. ahh!! wait, lets see an example
	arr := [5]int{1, 2, 3, 4, 5}
	sli := arr[:2] // copied first two indexs from arr
	arr[0] = 69    // it changes in sli too
	sli[1] = 96    // similarly changes in arr

	fmt.Printf("arr: %v, slice: %v, len = %d, cap = %d\n", arr, sli, len(sli), cap(sli)) // arr: [69 96 3 4 5], slice: [69 96], len = 2, cap = 5

	sli = append(sli, 6, 7) // apped 2 elements, so capacity will not increases
	arr[1] = 333
	sli[0] = 999
	fmt.Printf("arr: %v, slice: %v, len = %d, cap = %d\n", arr, sli, len(sli), cap(sli)) // arr: [999 333 6 7 5], slice: [999 333 6 7], len = 4, cap = 5

	sli = append(sli, 8, 9) // now what happen? len(sli) cross the cap(sli)
	arr[1] = 1
	sli[0] = 2
	fmt.Printf("arr: %v, slice: %v, len = %d, cap = %d\n", arr, sli, len(sli), cap(sli)) // arr: [999 1 6 7 5], slice: [2 333 6 7 8 9], len = 6, cap = 10

	// thats mean, slice copy array memory address and work as well until it does not crosses the capacity of array length
	// after crossing the array length slice create its own memory and be separeted from the array

	// Remove an element from slice
	s1 := []int{1, 2, 3, 4, 5}      // want to remove 3rd index (3)
	s2 := append(s1[:2], s1[3:]...) // need to make variadic as append() defined
	fmt.Println(s2)                 // [1 2 4 5]

	// MAP (just look a bit)
	m := map[string]int{
		"abc": 10,
		"xyz": 20,
	}
	fmt.Println(m, m["abc"]) // map[abc:10 xyz:20] 10

	val, ok := m["pqrs"]
	fmt.Println(val, ok) // 0 false
	m["pqrs"] = 30
	fmt.Println(m["pqrs"], val, ok) // 30 0 false ; here does not change

	delete(m, "pqrs") // simply will be deleted

	// try to play by making some function (higher-order)

	// for no reason - 1
	sqr := makeSquare([]int{1, 2, 3}, func(x int) int {
		return x * x
	})
	fmt.Println(sqr) // [1 4 9]

	// for no reason - 2
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evens := numFilter(nums, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println(evens) // [2 4 6 8]

	// for no reason - 3
	fmt.Println(people) // [{Cabul 10} {Abul 20} {Ebul 30} {Dabul 40} {Babul 50}]
	ple := whoIsYounger(people, func(x, y Person) bool {
		return x.Age < y.Age
	})
	fmt.Println(ple) // {Cabul 10}
}

func makeSquare(sli []int, sq func(x int) int) []int {
	res := make([]int, len(sli))
	for i, xx := range sli {
		res[i] = sq(xx) // HAHA!! we could 'xx*xx' instead
	}
	return res
}

func numFilter(sli []int, isEven func(x int) bool) []int {
	evs := make([]int, 0)
	for _, x := range sli {
		if isEven(x) {
			evs = append(evs, x)
		}
	}
	return evs
}

func whoIsYounger(people []Person, isChuto func(x, y Person) bool) Person {
	chuto := people[0]
	for _, p := range people {
		if isChuto(p, chuto) {
			chuto = p
		}
	}
	return chuto
}
