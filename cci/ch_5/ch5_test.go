package main

import (
	"fmt"
	"testing"
)

//\\//\\ 5.1 //\\//\\
func TestInsert(t *testing.T) {
	type insertTest struct {
		N             uint32
		M             uint32
		i             uint32
		j             uint32
		expectedValue uint32
	}

	fmt.Println("Testing Insert(N, M, i, j)")
	insertTests := []insertTest{
		{
			// 1011 with 100 at 2-0 --> 1100 = 8 + 4 = 12
			N:             11,
			M:             4,
			i:             0,
			j:             2,
			expectedValue: 12,
		},
		{
			// 1011 with 10 at 1-0 --> 1010 = 8+2 = 10
			N:             11,
			M:             2,
			i:             0,
			j:             1,
			expectedValue: 10,
		},
		{
			// 1011 with 10 at 1-2 --> 1101 = 8+4+1 = 13
			N:             11,
			M:             2,
			i:             1,
			j:             2,
			expectedValue: 13,
		},
		{
			// 100001 with 101 at 2-5 --> 110101 = 32+16+5 = 53
			N:             33,
			M:             5,
			i:             2,
			j:             4,
			expectedValue: 53,
		},
	}
	for _, insertTest := range insertTests {
		if val := Insert(insertTest.N, insertTest.M, insertTest.i, insertTest.j); val != insertTest.expectedValue {
			t.Fatalf("Insert(%d, %d, %d, %d) == val != expectdValue of %d", insertTest.N, insertTest.M, insertTest.i, insertTest.j, val, insertTest.expectedValue)
		}
	}
}

//\\//\\ 5.2 //\\//\\

//\\//\\ 5.3 //\\//\\
type bitsTest struct {
	N                    uint32
	expectedNextSmallest uint32
	expectedNextLargest  uint32
}

func bTest(n, ns, nl uint32) bitsTest {
	return bitsTest{
		N:                    n,
		expectedNextSmallest: ns,
		expectedNextLargest:  nl,
	}
}
func TestNexts(t *testing.T) {
	bitsTests := []bitsTest{
		bTest(0, 0, 0),
		bTest(3, 3, 5),
		bTest(1, 1, 2),
		bTest(2, 1, 4),
		bTest(15, 15, 23),
		bTest(1<<5, 1<<4, 1<<6),
		bTest(33, 17, 34),
	}
	fmt.Println("Testing Nexts(n)")
	for _, bTest := range bitsTests {
		if ns, nl := Nexts(bTest.N); ns != bTest.expectedNextSmallest && nl != bTest.expectedNextLargest {
			t.Fatalf("Nexts(%d) == (%d, %d), != expected values of (%d, %d)", bTest.N, ns, nl, bTest.expectedNextSmallest, bTest.expectedNextLargest)
		}
	}
}

//\\//\\ 5.4 //\\//\\
func TestIsPowerOfTwo(t *testing.T) {
	type powTwoTest struct {
		n  uint32
		is bool
	}
	pTtests := []powTwoTest{
		powTwoTest{0, true},
		powTwoTest{1, true},
		powTwoTest{2, true},
		powTwoTest{24, false},
		powTwoTest{100, false},
		powTwoTest{16, true},
	}
	fmt.Println("Testing IsPowerOfTwo(n)")
	for _, pTtest := range pTtests {
		if result := IsPowerOfTwo(pTtest.n); result != pTtest.is {
			t.Fatalf("IsPowerOfTwo(%d) == %t, != %t", pTtest.n, result, pTtest.is)
		}
	}
}

//\\//\\ 5.5. //\\//\\
func TestBitsShift(t *testing.T) {
	type bShiftTest struct {
		n        uint32
		m        uint32
		expected int
	}
	bShiftTests := []bShiftTest{
		{0, 1, 1},
		{1, 0, 1},
		{31, 14, 2},
	}
	fmt.Println("Testing BitsToShift(n, m)")
	for _, bst := range bShiftTests {
		if result := BitsToShift(bst.n, bst.m); result != bst.expected {
			t.Fatalf("BitsToShift(%d, %d) == %d, != %d", bst.n, bst.m, result, bst.expected)
		}
	}
	fmt.Println("Testing BitsToShiftXOR(n, m)")
	for _, bst := range bShiftTests {
		if result := BitsToShiftXOR(bst.n, bst.m); result != bst.expected {
			t.Fatalf("BitsToShift(%d, %d) == %d, != %d", bst.n, bst.m, result, bst.expected)
		}
	}
}

//\\//\\ 5.6 //\\//\\
func TestSwapEvenAndOddBits(t *testing.T) {
	type swapTest struct {
		n        uint32
		expected uint32
	}
	swapTests := []swapTest{
		{0xA, 0x5},
		{0x0, 0x0},
		{0x2, 0x1},
		{0xAC, 0x5C},
		{0xF, 0xF},
	}
	fmt.Println("Testing SwapEvenAndOddBits(n)")
	for _, st := range swapTests {
		if result := SwapEvenAndOddBits(st.n); result != st.expected {
			t.Fatalf("SwapEvenAndOddBits(%d) == %d, != %d", st.n, result, st.expected)
		}
	}
}

//\\//\\ 5.8 //\\//\\
func TestDrawHorizontalLine(t *testing.T) {
	type hLineTest struct {
		screen         []byte
		width          int
		x1             int
		x2             int
		y              int
		expectedScreen []byte
	}
	hLineTests := []hLineTest{
		{
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			3,
			0,
			0,
			0,
			[]byte{0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			3,
			0,
			2,
			1,
			[]byte{0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00},
		},
		{
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			3,
			0,
			1,
			1,
			[]byte{0x00, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00},
		},
		{
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			3,
			0,
			1,
			4,
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			3,
			4,
			1,
			2,
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			3,
			0,
			4,
			1,
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
	}
	fmt.Println("Testing DrawHorizontalLine(screen, width, x1, x2, y)")
	for _, hlt := range hLineTests {
		if result := DrawHorizontalLine(hlt.screen, hlt.width, hlt.x1, hlt.x2, hlt.y); !equalScreens(result, hlt.expectedScreen) {
			fmt.Println(result)
			t.Fatalf("DrawHorizontalLine(%v, %d, %d, %d, %d) == %v, != %v", hlt.screen, hlt.width, hlt.x1, hlt.x2, hlt.y, result, hlt.expectedScreen)
		}
	}
}

func equalScreens(screen1, screen2 []byte) bool {
	if len(screen1) != len(screen2) {
		return false
	}
	for i := range screen1 {
		if screen1[i] != screen2[i] {
			return false
		}
	}
	return true
}
