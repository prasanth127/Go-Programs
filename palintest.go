//go: palindrome_test.go (automated unit test program for testing palindrome functionality)

//package word
package main

import "unicode"

//import "fmt"

func IsPalindrome1(s string) bool {
	for i := range s { //a2a//len =3, start_index = 0; last_inde = 2
		if s[i] != s[len(s)-1-i] { //i=0;s[0] == s[2] //s[1] =2; s[1] = 2
			return false
		}
	}
	return true
}
func IsPalindrome(s string) bool {
	var letters []rune //rune is a UTF8 unicode character in golang
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
