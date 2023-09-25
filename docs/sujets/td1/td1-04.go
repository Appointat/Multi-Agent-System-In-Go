package main

import (
	"fmt"
)

dict := [...]string{"AGENT", "CHIEN", "COLOC", "ETANG", "ELLE", "GEANT", "NICHE", "RADAR"}

func isPalindrome (word string) bool {
	l := len(word)
	for i := 0; i < l/2; i++ {
		if word[i] != word[l-i-1] {
			return false
		}
	}
	return true
}

func Palindromes (words []string) (l []string) {
	l := make([]string, 0)
	for word := range words {
		if isPalindrome(word) {
			l = append(l, word)
		}
	}
}

func main() {
	fmt.Println(Palindromes(dict))
}