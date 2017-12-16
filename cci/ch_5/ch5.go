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

func clearBitsItoZero(num, i uint) uint {
	return num & ^((1 << (i + 1)) - 1)
}

func clearBitsMSBtoI(num, i uint) uint {
	return num & ((1 << i) - 1)
}

//\\//\\ 5.1 //\\//\\
// Assumes M will be fit within N between j and i.
// j and i have at least enough space for M.
func Insert(N, M, i, j uint) uint {
	// Find the section N between j and i
	// clear this section
	// or this section with M shift by (j - i)
	return 1
}

//\\//\\//\\ MAIN //\\//\\//\\
func main() {
	fmt.Println(problemTitle("5.1"))
	fmt.Println(getBit(2, 1) == 1)
	fmt.Println(setBit(5, 1) == 7)
	fmt.Println(clearBit(7, 1) == 5)
	fmt.Println(clearBitsItoZero(15, 2) == 8)
	fmt.Println(clearBitsMSBtoI(15, 2) == 3)
	fmt.Println(problemEndline(problemTitle("5.1")))
}
