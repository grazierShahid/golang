package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	num := 42
	text := "Hello"
	isTrue := true
	fmt.Printf("Default: %v, Type: %T, Binary: %b, Decimal: %d, Octal: %o, Hex: %x\n", num, num, num, num, num, num)
	fmt.Printf("String: %s, Quoted: %q, Float: %.2f, Boolean: %t, Pointer: %p\n", text, text, 3.14159, isTrue, &num)

	// set the bufio to read the user standard input continiously
	reader := bufio.NewReader(os.Stdin)

	// assume, we are going to take a name and age as input
	input, _ := reader.ReadString('\n')
	fields := strings.Fields(input)

	// extract the name and age
	name := strings.Join(fields[:len(fields)-1], " ")
	age, _ := strconv.Atoi(fields[len(fields)-1]) // Atoi -> alphabet to int
	str := fmt.Sprintf("Hey %s! you are %d years old!\n", name, age)

	// now we want to store this str into a file, so let's create a text file
	file, _ := os.Create("Output.txt")
	fmt.Fprint(file, str) // it will save the str to Output.txt

	// NOW LET'S GO FOR STRING

	// how string store byte(as char[ASCII]) as like array
	str = "grazier"
	for ind, char := range str {
		fmt.Printf("index: %d, char: %c, byte: %x\n", ind, char, char)
	}
	/*
		output:
		index: 0, char: g, byte: 67
		index: 1, char: r, byte: 72
		index: 2, char: a, byte: 61
		index: 3, char: z, byte: 7a
		index: 4, char: i, byte: 69
		index: 5, char: e, byte: 65
		index: 6, char: r, byte: 72
	*/

	//If we want to use rune(UNICODE) instead of byte(ASCII), we need to convert
	runes := []rune(str)
	for i, r := range runes {
		fmt.Printf("index: %d, char: %c, Unicode: %U, hex: %x\n", i, r, r, r)
	}
	/*
		output:
		index: 0, char: g, Unicode: U+0067, hex: 67
		index: 1, char: r, Unicode: U+0072, hex: 72
		index: 2, char: a, Unicode: U+0061, hex: 61
		index: 3, char: z, Unicode: U+007A, hex: 7a
		index: 4, char: i, Unicode: U+0069, hex: 69
		index: 5, char: e, Unicode: U+0065, hex: 65
		index: 6, char: r, Unicode: U+0072, hex: 72
	*/

	// at chapter 6 i wrote a lot about rune vs byte vs string
	// so now time to see string alone

	// string comparison - we can compare string as like int (Lexicographical comparison)
	str1, str2 := "abc", "xyz"
	fmt.Println(str1 == str2, str1 < str2) // false ture
	str2 = "abc"
	fmt.Println(str1 == str2) // true

	// we can assign the comparison result in variable too
	bool1 := str1 == str2
	bool2 := strings.Compare(str1, str2)
	fmt.Println(bool1, bool2) // true 0

	// concatenating using += operator
	str1 += str2
	str = strings.Join([]string{str1, str2}, "-") // the last "-" will set in the middle at both string
	fmt.Println(str1, str)                        // abcabc abcabc-abc

	// trim string
	str = "@#Hell0@#World@#"
	fmt.Println(strings.Trim(str, "@#"))         // Hell0@#World
	fmt.Println(strings.TrimLeft(str, "@#"))     // Hell0@#World@#
	fmt.Println(strings.TrimRight(str, "@#"))    // @#Hell0@#World
	fmt.Println(strings.TrimSpace("  a b  c  ")) // a b  c
	fmt.Println(strings.TrimPrefix(str, "@#"))   // Hell0@#World@#
	fmt.Println(strings.TrimSuffix(str, "@#"))   // @#Hell0@#World

	// split string
	str = "welcome,to,grazier's,github"
	fmt.Println(strings.Split(str, ","))          // [welcome to grazier's github]
	fmt.Println(strings.SplitAfter(str, ","))     // [welcome, to, grazier's, github]
	fmt.Println(strings.SplitN(str, ",", 4))      // [welcome to grazier's github]
	fmt.Println(strings.SplitAfterN(str, ",", 2)) // [welcome, to,grazier's,github]

	// check that is there any substring acontain in another string; return bool value
	str1, str2 = "grazierShahid", "zier"
	fmt.Println(strings.Contains(str1, str2))     // true
	fmt.Println(strings.ContainsAny(str1, "abc")) // true
	fmt.Println(strings.Contains(str1, "abc"))    // false

	// if we want to repeat a string, value should be positive
	str = "Grazier"
	fmt.Println(strings.Repeat(str, 3)) // GrazierGrazierGrazier

	// to find a substring's first index and count the num of substr
	str = "His name is Abul, Abul is a brother of Babul"
	fmt.Println(strings.Index(str, "Abul")) // 12
	fmt.Println(strings.Count(str, "Abul")) // 2\
}
