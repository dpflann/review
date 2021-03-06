package main

import (
	"bytes"
	"fmt"
	"math"
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

func getBit(num uint32, i uint32) uint32 {
	if (num & (1 << i)) == 0 {
		return 0
	}
	return 1
}

func setBit(num, i uint32) uint32 {
	return num | (1 << i)
}

func clearBit(num, i uint32) uint32 {
	return num & ^(1 << i)
}

func clearBitsItoZero(num, i uint32) uint32 {
	return num & ^((1 << (i + 1)) - 1)
}

func clearBitsMSBtoI(num, i uint32) uint32 {
	return num & ((1 << i) - 1)
}

//\\//\\ 5.1 //\\//\\
// Assumes M will be fit within N between j and i.
// j and i have at least enough space for M.
func Insert(N, M, i, j uint32) uint32 {
	// Find the section N between j and i
	// clear this section
	// or this section with M shift by (j - i)
	var leftSide, rightSide uint32
	if j+1 > 31 {
		leftSide = N
	} else {
		leftSide = clearBitsMSBtoI(N, j+1)
	}
	if i-1 < 0 {
		rightSide = N
	} else {
		rightSide = clearBitsItoZero(N, i-1)
	}
	mask := leftSide ^ rightSide
	return mask | (M << i)
}

//\\//\\ 5.2 //\\//\\

//\\//\\ 5.3 //\\//\\
// for N with n bits set to 1, find the next smallest and next largest with m = n bits set to 1
func countOnes(n uint32) int {
	count := 0
	for n > 0 {
		if (n & 1) == 1 {
			count += 1
		}
		n = n >> 1
	}
	return count
}

func Nexts(n uint32) (uint32, uint32) {
	if n == 0 || n == (1<<32-1) {
		return n, n
	}
	ones := countOnes(n)
	if ones == 1 {
		return n >> 1, n << 1
	}
	ns, nl := n-1, n+1
	for ns > 0 && countOnes(ns) != ones {
		ns -= 1
	}
	for nl < (1<<32-1) && countOnes(nl) != ones {
		nl += 1
	}
	return ns, nl
}

//\\//\\ 5.4 //\\//\\
func IsPowerOfTwo(n uint32) bool {
	return (n & (n - 1)) == 0
}

//\\//\\ 5.5 //\\//\\
func BitsToShift(n, m uint32) int {
	nOnes := countOnes(n)
	mOnes := countOnes(m)
	// n to m
	return int(math.Abs(float64(nOnes - mOnes)))
}

func BitsToShiftXOR(n, m uint32) int {
	difference := n ^ m
	return countOnes(difference)
}

//\\//\\ 5.6 //\\//\\
func SwapEvenAndOddBits(n uint32) uint32 {
	var oddMask uint32 = 0xAAAAAAAA
	return ((n&oddMask)>>1 | (n & ^oddMask)<<1)
}

//\\//\\ 5.7 //\\//\\

//\\//\\ 5.8 //\\//\\
func DrawHorizontalLine(screen []byte, width, x1, x2, y int) []byte {
	if x1 > (width-1) || x2 > (width-1) || (y > (len(screen) / width)) {
		return screen
	}
	r := y * width
	c1 := x1 + r
	c2 := x2 + r
	for i := range screen[c1 : c2+1] {
		screen[c1+i] = 0xFF
		i += 1
	}
	return screen
}

//\\//\\//\\ MAIN //\\//\\//\\
func main() {
	fmt.Println(problemTitle("5.1"))
	fmt.Println(getBit(2, 1) == 1)
	fmt.Println(setBit(5, 1) == 7)
	fmt.Println(clearBit(7, 1) == 5)
	fmt.Println(clearBitsItoZero(15, 2) == 8)
	fmt.Println(clearBitsItoZero(15, 0) == 14)
	fmt.Println(clearBitsItoZero(15, 0) == 14)
	fmt.Println(clearBitsMSBtoI(15, 2) == 3)
	fmt.Println("Insert(N, M, i, j)")
	// 1011 with 100 at 2-0 --> 1100 = 8 + 4 = 12
	fmt.Println(Insert(11, 4, 0, 2) == 12)
	// 1011 with 10 at 1-0 --> 1010 = 8+2 = 10
	fmt.Println(Insert(11, 2, 0, 1) == 10)
	// 1011 with 10 at 1-2 --> 1101 = 8+4+1 = 13
	fmt.Println(Insert(11, 2, 1, 2) == 13)
	// 100001 with 101 at 2-5 --> 110101 = 32+16+5 = 53
	fmt.Println(Insert(33, 5, 2, 4) == 53)
	fmt.Println(problemEndline(problemTitle("5.1")))
	// All tests are now in ch5_test.go
}
