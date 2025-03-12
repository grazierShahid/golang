# Working with Strings and I/O in Go

To start with string, we need to learn byte, rune and some i/o method. As well we already learned about byte and rune in chapter 6, so now only remain about i/o methods. So lets start by bufio package.

The bufio package implements buffered I/O operations that make input and output more efficient. It wraps an io.Reader or io.Writer object, creating another object that also implements the interface but provides buffering and some help for textual I/O.

## Terminal I/O Methods

Using `fmt`, we can take input using `Scan()`, `Scanf()`, and `Scanln()`. Then we can print output using `Print()`, `Printf()`, and `Println()`. Lets see why there are diff func for the same thing:

**Input Functions:**

- `Scan()`: Reads space-separated values into provided variables
- `Scanf()`: Reads input according to format specifiers for precise control
- `Scanln()`: Like Scan, but stops at newline

**Output Functions:**

- `Print()`: Outputs values without a newline
- `Printf()`: Outputs formatted text using specifiers
- `Println()`: Outputs values followed by a newline

## File I/O Methods

All avobe function are for read and write in terminal. But if we want to make file in a codebase and read and write with there, then we need to take help from `bufio` and `os` packages.

The `NewReader()` function from the `bufio` package can help to read input from user with help of `os.Stdin`. See the below code:

```go
reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')
```

The `reader` variable can help to read input from user continiously. `ReadString()` function read the `string` given by user, and it takes as input until new line arrive (for `'\n'`). After that, we need to separete the `input` using `strings.Fields()`. See the code below:

```go
fields := strings.Fields(input)
```

The `Fields()` function splits the input string around whitespace (spaces, tabs, newlines) and returns a slice containing the non-empty substrings.

So, now if we need to work with the `fields` valiable, we can do anything as we need.

## How to create file and work with this

The `os.Create()` helps to create a file, and `fmt.Fprint()` helps to write something in this file. So let's see the given code:

```go
file, err := os.Create("output.txt")
fmt.Fprint(file, str) // save str to file/output.txt
```

So, now we are able to read and write with file system!

## Built-it function in strings package

String is a data structure which store each character, and it is immutable in golang. And `strings` package have lot of function to work with string.

- All the function to trim a string:

```go
func Trim(s string, cutset string) string
func TrimLeft(s string, cutset string) string
func TrimRight(s string, cutset string) string
func TrimSpace(s string) string
func TrimPrefix(s, prefix string) string
func TrimSuffix(s, suffix string) string
```

- All the function for string split:

```go
func Split(s, sep string) []string
func SplitAfter(s, sep string) []string
func SplitN(s, sep string, n int) []string
func SplitAfterN(str, sep string, n int) []string
// n for both SplitN and SplitAfterN
//n > 0: Splits up to n substrings.
//n == 0: Returns an empty slice.
//n < 0: Returns all substrings.
```

And there are lot of function like Contains(), ContainsAny(), Repeat(), Index().

Those are the basic needy func from the `strings` package.
