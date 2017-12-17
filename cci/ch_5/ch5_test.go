package main

import (
	"fmt"
	"testing"
)

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
			t.Fatal("Insert(%d, %d, %d, %d) == val != expectdValue of %d", insertTest.N, insertTest.M, insertTest.i, insertTest.j, val, insertTest.expectedValue)
		}
	}
}
