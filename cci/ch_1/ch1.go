package main

import (
	"bytes"
	"fmt"
	"strconv"
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

// ====== 1.4 ====== \\
// Attempt to work with  []rune because of the idea of 'true' length
func ReplaceSpaces(s *string) {
	if s == nil {
		return
	}
	spaces := 0
	for _, c := range *s {
		if c == ' ' {
			spaces++
		}
	}
	if spaces < 1 {
		return
	}
	newStringBuff := make([]rune, len(*s)+spaces*2) // adding 2 more chars for each space
	i := 0
	for _, c := range *s {
		if c == ' ' {
			newStringBuff[i] = '%'
			newStringBuff[i+1] = '2'
			newStringBuff[i+2] = '0'
			i += 3
		} else {
			newStringBuff[i] = c
			i += 1
		}
	}

	newS := string(newStringBuff)
	*s = newS
}

// The solution in the text assumes null terminated char array.
// This can be faked, but Go does not work with strings that way.
func ReplaceSpacesCCI(str *[]rune, length int) {
	spaceCount, newLength, i := 0, 0, 0
	for i < length {
		if (*str)[i] == ' ' {
			spaceCount++
		}
		i++
	}
	newLength = length + spaceCount*2
	for i = length - 1; i >= 0; i-- {
		if (*str)[i] == ' ' {
			(*str)[newLength-1] = '0'
			(*str)[newLength-2] = '2'
			(*str)[newLength-3] = '%'
			newLength = newLength - 3
		} else {
			(*str)[newLength-1] = (*str)[i]
			newLength = newLength - 1
		}
	}
}

// ====== 1.5 ====== \\
func Compress(s string) string {
	lenS := len(s)
	if lenS == 0 {
		return s
	}
	Srunes := []rune(s)
	compressedS := ""
	currentChar := Srunes[0]
	count := 1
	for i := 1; i < lenS; i++ {
		if Srunes[i] == currentChar {
			count += 1
		} else {
			compressedS += string(currentChar) + strconv.Itoa(count)
			currentChar = Srunes[i]
			count = 1
		}
	}
	// final step because the for-loop looks ahead
	compressedS += string(currentChar) + strconv.Itoa(count)
	if len(compressedS) > lenS {
		return s
	}
	return compressedS
}

func CompressCCI(s string) string {
	lenS := len(s)
	if lenS == 0 {
		return s
	}
	var buffer bytes.Buffer
	currentChar := s[0]
	count := 1
	for i := 1; i < lenS; i++ {
		if s[i] == currentChar {
			count += 1
		} else {
			buffer.WriteByte(currentChar)
			buffer.WriteString(strconv.Itoa(count))
			currentChar = s[i]
			count = 1
		}
	}
	// final step because the for-loop looks ahead
	buffer.WriteByte(currentChar)
	buffer.WriteString(strconv.Itoa(count))
	compressedS := buffer.String()
	if len(compressedS) > lenS {
		return s
	}
	return compressedS

}

func main() {
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

	fmt.Println("====== 1.4 ======")
	spaces1 := "a b"
	spaces2 := "a  b "
	expectedSpaces1 := "a%20b"
	expectedSpaces2 := "a%20%20b%20"
	noSpaces := "ab"
	fmt.Println("ReplaceSpaces(...)")
	ReplaceSpaces(&spaces1)
	fmt.Println(spaces1 == expectedSpaces1)
	ReplaceSpaces(&noSpaces)
	fmt.Println(noSpaces == noSpaces)
	ReplaceSpaces(&spaces2)
	fmt.Println(spaces2 == expectedSpaces2)
	fmt.Println("ReplaceSpacesCCI(...)")
	runeSpaces1 := []rune{'a', ' ', 'b', ' ', ' '}
	runeSpaces1Expected := []rune{'a', '%', '2', '0', 'b'}
	lenRuneSpaces1 := 3
	ReplaceSpacesCCI(&runeSpaces1, lenRuneSpaces1)
	fmt.Println(string(runeSpaces1) == string(runeSpaces1Expected))

	fmt.Println("====== 1.5 ======")
	s1Uncompressed := "abc"
	s1Compressed := "abc"
	fmt.Println("Compress(...)")
	fmt.Println(Compress(s1Uncompressed) == s1Compressed)
	s2Uncompressed := "a"
	s2Compressed := "a"
	fmt.Println(Compress(s2Uncompressed) == s2Compressed)
	s3Uncompressed := "aaabbbccc"
	s3Compressed := "a3b3c3"
	fmt.Println(Compress(s3Uncompressed) == s3Compressed)
	fmt.Println(Compress("") == "")
	fmt.Println("CompressCCI(...)")
	fmt.Println(CompressCCI(s1Uncompressed) == s1Compressed)
	fmt.Println(CompressCCI(s2Uncompressed) == s2Compressed)
	fmt.Println(CompressCCI(s3Uncompressed) == s3Compressed)
	fmt.Println(CompressCCI("") == "")
}
