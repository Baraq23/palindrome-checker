# Palindrome Checker

This Go program checks whether a given word is a palindrome.

## Description

The program takes a single command-line argument (a word) and determines if it is a palindrome. A palindrome is a word that reads the same backward as forward, case insensitive.

## Installation

To run this program, you need to have Go installed on your machine. You can download and install Go from the [official website](https://golang.org/dl/).

## Usage

1. Clone the Palindrome-checker repo [here](https://github.com/Baraq23/palindrome-checker.git)
2. Open your terminal.
3. Navigate to the directory where `main.go` is saved.
4. Run the program with the command:
   ```bash
   go run . <word>
   ```
   Replace `<word>` with the word you want to check.

Example:
```bash
go run . radar
```

## Functions

### ValidateWord

```go
func ValidateWord(s string)
```
- Validates that the input string consists of alphabetic characters only.
- Prints an error message if the input contains non-alphabetic characters.

### Palindrome

```go
func Palindrome(s string) bool
```
- Takes a string as input and returns `true` if the string is a palindrome, `false` otherwise.
- Converts uppercase characters to lowercase for case insensitivity.

## Example

Running the command:
```bash
go run . Radar
```
Will output:
```
Radar is a palindrome.
```

Running the command:
```bash
go run . Hello
```
Will output:
```
Hello is NOT a palindrome.
```

## License

This project is licensed under the MIT License.