package main

import (
	"bytes"
	"fmt"
)

func problemTitle(s string) string {
	side := "====="
	return fmt.Sprintf("%s %s %s", side, s, side)
}

func problemEndline(t string) string {
	b := bytes.Buffer{}
	for range t {
		b.WriteString("=")
	}
	return b.String()
}

func getBit(num uint, i uint) uint {
	if (num & (1 << i)) == 0 {
		return 0
	}
	return 1
}

func setBit(num, i uint) uint {
	return num | (1 << i)
}

func clearBit(num, i uint) uint {
	return num & ^(1 << i)
}

//\\//\\//\\ MAIN //\\//\\//\\
func main() {
	fmt.Println(problemTitle("5.1"))
	fmt.Println(getBit(2, 1) == 1)
	fmt.Println(setBit(5, 1) == 7)
	fmt.Println(clearBit(7, 1) == 5)
	fmt.Println(problemEndline(problemTitle("5.1")))
}
