package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <word>")
		return
	}
	word := os.Args[1]
	ValidateWord(word)
	palindrome := Palindrome(word)

	if palindrome {
		fmt.Printf("%v is a palindrome.\n", word)
	} else {
		fmt.Printf("%v is NOT a palindrome.\n", word)
	}
}

// Function ValidateWord() checks and validates the word given
func ValidateWord(s string) {
	for _, chr := range s {
		if (chr >= 'a' && chr <= 'z') || (chr >= 'A' && chr <= 'Z') {
			continue
		} else {
			fmt.Printf("%v is not a word\n", s)
			os.Exit(0)
		}
	}
}

// Function Palindrome() takes a word then returns its palindrome
func Palindrome(s string) bool {
	str := ""
	word := ""

	for _, l := range s {
		if l >= 'A' && l <= 'Z' {
			l += 32
			str += string(l)
			word = string(l) + word

		} else {
			str += string(l)
			word = string(l) + word
		}
	}
	return str == word
}
