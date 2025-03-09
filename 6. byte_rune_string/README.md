# Understanding Runes and Bytes in Go

## Introduction

This chapter is about the pre-requisite of string, and the pre-requisites are byte and rune in golang. Because in golang, you cannot change a character in a string. Either you can change the whole value of a string, or you should convert the string into a `[]byte` or `[]rune`.

So, let's start by introducing what are the byte and rune in golang.

## Bytes and Runes - The Building Blocks

In Go, strings are immutable sequences of bytes. This design choice significantly affects how we work with text in Go, especially when dealing with international text and Unicode characters.

## What is a Byte?

A byte in Go is represented by the type `byte`, which is an alias for `uint8` (an 8-bit unsigned integer). It can represent 256 different values (0-255). For ASCII characters, a single byte is sufficient to represent one character. However, many Unicode characters require multiple bytes.

## What is a Rune?

A rune in Go is represented by the type `rune`, which is an alias for `int32` (a 32-bit signed integer). It can represent a Unicode code point, which is essentially any character in the Unicode standard. This allows Go to handle any character from any language, including emojis and special symbols.

## String Representation

When working with strings in Go, it's important to understand that:

1. Strings are stored as sequences of bytes internally
2. String literals are UTF-8 encoded by default
3. The `len()` function returns the number of bytes, not characters

## String-Byte-Rune Conversions

Go provides straightforward conversion between strings, bytes, and runes:

```go
// String to byte slice
str := "grazier"
bytes := []byte(str)

// Byte slice to string
str2 := string(bytes)

// String to rune slice
runes := []rune(str)

// Rune slice to string
str3 := string(runes)
```

## String Iteration

There are two main ways to iterate over strings in Go:

### Using Regular for Loop (Iterates Over Bytes)

```go
str := "abc"
for i := 0; i < len(str); i++ {
    fmt.Printf("index: %d, char: %v, value %v, type: %T\n", i, string(str[i]), str[i], str[i])
}
```

Output:

```
index: 0, char: a, value 97, type: uint8
index: 1, char: b, value 98, type: uint8
index: 2, char: c, value 99, type: uint8
```

When using a `for range` loop, each iteration decodes one UTF-8-encoded rune, giving you the proper Unicode characters regardless of how many bytes they occupy.

## Character Manipulation

Since strings are immutable in Go, you cannot directly modify a character in a string by index. Here are different approaches to modify characters:

## Using String Slicing

```go
str := "abc"
str = str[:1] + "x" + str[2:]  // Replace the character at index 1
// Results in "axc"
```

## Using Byte Slice

```go
str := "abc"
bytes := []byte(str)
bytes[1] = 'B'  // Replace the character at index 1
str = string(bytes)
// Results in "aBc"
```

## Using Rune Slice

```go
str := "abc"
runes := []rune(str)
runes[1] = 'X'  // Replace the character at index 1
str = string(runes)
// Results in "aXc"
```

## Unicode Support

Go has excellent Unicode support, which is particularly important for handling non-ASCII text like Bengali, Chinese, or even emojis.

## Example with Bengali Text

```go
s := "বাংলাদেশ - 🤸🏾"

// Converting to bytes and runes
withByte := []byte(s)
withRune := []rune(s)

// Length comparison
fmt.Printf("String length (bytes): %d\n", len(s))         // 33
fmt.Printf("Rune slice length (characters): %d\n", len([]rune(s))) // 13
```

This example demonstrates:

1. UTF-8 encoding in action
2. The difference between byte representation (each Unicode character takes multiple bytes)
3. Rune representation (each Unicode character is a single rune)
4. The difference in length when counting bytes versus characters

## Understanding Length Differences

```go
text := "hello, সোনার ছেলে!"
fmt.Printf("%d\n", len(text))         // 36 (bytes)
fmt.Printf("%d\n", len([]rune(text))) // 18 (characters)
fmt.Printf("%d\n", len([]byte(text))) // 36 (bytes)
```

This clearly shows that the string "hello, সোনার ছেলে!" contains:

- 36 bytes (due to multi-byte Unicode characters)
- 18 actual characters (runes)

## Creating Strings from Rune Slices

You can directly create strings from rune slices:

```go
R := []rune{'অ', 'আ', 'ক', 'খ'}
S := string(R)
fmt.Println(S)  // Outputs: অআকখ
```

## Why This Matters

Understanding the difference between bytes and runes is crucial in Go because:

1. It affects how you iterate through strings
2. It impacts how you measure string length
3. It determines how you manipulate individual characters
4. It's essential for proper handling of international text

By mastering these concepts, you'll be able to handle text processing in Go with confidence, regardless of the language or character set you're working with.

## Key Takeaways

1. Strings in Go are immutable sequences of bytes
2. A byte (`uint8`) can represent 256 different values
3. A rune (`int32`) can represent any Unicode code point
4. Regular `for` loops iterate over bytes, while `for range` loops iterate over runes
5. To modify characters in a string, convert it to a byte or rune slice first
6. The `len()` function returns the number of bytes, not characters
7. For proper character counting of Unicode text, convert to runes first
