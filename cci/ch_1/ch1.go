package main

import (
	"fmt"
)

// ====== 1.1 ====== \\
func HasUnique(s string) bool {
	ls := len(s)
	for current := 0; current < ls; current++ {
		c := s[current]
		for i := current + 1; i < ls; i++ {
			if c == s[i] {
				return false
			}
		}
	}
	return true
}

func HasUniqueDS(s string) bool {
	charExists := make(map[byte]bool, len(s))
	for i := range s {
		if _, ok := charExists[s[i]]; ok {
			return false
		}
		charExists[s[i]] = true
	}
	return true
}

// Assumes character set is ASCII
func HasUniqueCCI(s string) bool {
	if len(s) > 256 {
		return false
	}
	r := []rune(s)
	usedChars := make([]bool, 256)
	for _, c := range r {
		if usedChars[c] {
			return false
		}
		usedChars[c] = true
	}
	return true
}

// ====== 1.2 ====== \\
func Reverse(str *string) {
	if str == nil {
		return
	}
	r := []rune(*str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	*str = string(r)
}

// ====== 1.3 ====== \\
func IsPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1Chars := make(map[rune]int, len(s1))
	for _, c := range s1 {
		s1Chars[c] = s1Chars[c] + 1
	}
	for _, c := range s2 {
		if count, ok := s1Chars[c]; !ok || count < 0 {
			return false
		}
		s1Chars[c] = s1Chars[c] - 1
	}
	return true
}

func IsPermutationCCI(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// Assume ASCII
	letters := make([]int, 256)
	for _, c := range s1 {
		letters[c]++
	}
	for _, c := range s2 {
		letters[c] -= 1
		if letters[c] < 0 {
			return false
		}
	}
	return true
}

func main() {
	// TODO: update to have shared variables
	fmt.Println("====== 1.1 ======")
	fmt.Println("HasUnique(...)")
	fmt.Println(HasUnique("abc") == true)
	fmt.Println(HasUnique("abca") == false)
	fmt.Println("HasUniqueDS(...)")
	fmt.Println(HasUniqueDS("abc") == true)
	fmt.Println(HasUniqueDS("abca") == false)
	fmt.Println("HasUniqueCCI(...)")
	fmt.Println(HasUniqueCCI("abc") == true)
	fmt.Println(HasUniqueCCI("abca") == false)
	fmt.Println("=================")

	fmt.Println("====== 1.2 ======")
	hello := "hello"
	reversed := "olleh"
	s := &hello
	Reverse(s)
	fmt.Println("Reverse(...)")
	fmt.Println(reversed == *s)

	fmt.Println("====== 1.3 ======")
	s1 := "abc"
	s2 := "cba"
	s3 := "ab"
	s4 := "aabc"
	fmt.Println("IsPermutation(..., ...)")
	fmt.Println(IsPermutation(s1, s1) == true)
	fmt.Println(IsPermutation(s1, s2) == true)
	fmt.Println(IsPermutation(s1, s3) == false)
	fmt.Println(IsPermutation(s1, s4) == false)
	fmt.Println("IsPermutationCCI(..., ...)")
	fmt.Println(IsPermutationCCI(s1, s1) == true)
	fmt.Println(IsPermutationCCI(s1, s2) == true)
	fmt.Println(IsPermutationCCI(s1, s3) == false)
	fmt.Println(IsPermutationCCI(s1, s4) == false)
}
