package main

import "fmt"

func main() {
	// string - byte conversation : byte -> uint8
	str := "grazier"
	bytes := []byte(str)
	fmt.Println(bytes) // [103 114 97 122 105 101 114]
	str2 := string(bytes)
	fmt.Println(str2) // grazier

	// string - rune conversation : rune -> int32
	str3 := "grazier"
	runes := []rune(str3)
	fmt.Println(runes) // [103 114 97 122 105 101 114]
	str4 := string(runes)
	fmt.Println(str4) // grazier

	// rune create by slice itself
	temp := []byte{0x67, 0x72, 0x61, 0x7a, 0x69, 0x65, 0x72}
	tempStr := string(temp)
	fmt.Println(temp)    // [103 114 97 122 105 101 114]
	fmt.Println(tempStr) // grazier

	str = "abc" // string is immutable (non-changable on index)

	// for loop (as the next line) hold string as byte
	for i := 0; i < len(str); i++ {
		fmt.Printf("index: %d, char: %v, value %v, type: %T\n", i, string(str[i]), str[i], str[i])
	}
	/* outputs:
	index: 0, char: a, value 97, type: uint8
	index: 1, char: b, value 98, type: uint8
	index: 2, char: c, value 99, type: uint8
	*/

	// for range loop decodes one UTF-8-encoded rune on each iteration
	for index, value := range str {
		fmt.Printf("index: %d, char: %v, value %v, type: %T\n", index, string(value), value, value)
	}
	/* outputs:
	index: 0, char: a, value 97, type: int32
	index: 1, char: b, value 98, type: int32
	index: 2, char: c, value 99, type: int32
	*/

	// convert rune to byte
	for index, value := range str {
		runeToByte := byte(value)
		fmt.Printf("index: %d, char: %v, value %v, type: %T\n", index, string(value), runeToByte, runeToByte)
	}
	/* outputs:
	index: 0, char: a, value 97, type: uint8
	index: 1, char: b, value 98, type: uint8
	index: 2, char: c, value 99, type: uint8
	*/

	// if you want to change a single char, then
	// str[1] = 'x' --> it gives error. instead of this you can
	str = str[:1] + "x" + str[2:]
	fmt.Println(str) // axc
	// or you can change using []byte or []rune

	// using []byte
	bytes = []byte(str)
	bytes[1] = 'B'
	str = string(bytes)
	fmt.Println(str)

	// using []rune
	runes = []rune(str) // aBc
	runes[1] = 'X'
	str = string(runes)
	fmt.Println(str) // aXc

	//
	s := "বাংলাদেশ - 🤸🏾"
	withByte := []byte(s)
	withRune := []rune(s)
	fmt.Printf("type: %T, value: %v\n", s, s)               // type: string, value: বাংলাদেশ - 🤸🏾
	fmt.Printf("type: %T, value: %v\n", withByte, withByte) // []uint8, value: [224 166 172 224 166 190 224 166 130 224 166 178 224 166 190 224 166 166 224 167 135 224 166 182 32 45 32 240 159 164 184 240 159 143 190]
	fmt.Printf("type: %T, value: %v\n", withRune, withRune) // []int32, value: [2476 2494 2434 2482 2494 2470 2503 2486 32 45 32 129336 127998]
	fmt.Println(string(withByte), string(withRune))         // বাংলাদেশ - 🤸🏾 বাংলাদেশ - 🤸🏾

	// length -- string vs byte vs rune
	text := "hello, সোনার ছেলে!"
	fmt.Printf("%d\n", len(text))         // 36
	fmt.Printf("%d\n", len([]rune(text))) // 18
	fmt.Printf("%d\n", len([]byte(text))) // 36

	// we can make string by rune slice directly
	R := []rune{'অ', 'আ', 'ক', 'খ'} // no more than one char
	S := string(R)
	fmt.Println(S) // অআকখ

}
