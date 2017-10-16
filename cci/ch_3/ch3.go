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

// ===== 3.1 =====

type ThreeStackArray struct {
	stack         []int
	stackPointers []int
	stackSegments [][]int
	stackSize     int
}

func NewThreeStackArray(size int) *ThreeStackArray {
	t := ThreeStackArray{
		stack:     make([]int, 3*size),
		stackSize: size,
	}
	stackSegments := make([][]int, 3)
	s1 := 0
	e1 := size - 1
	s2 := size
	e2 := size + size - 1
	s3 := size + size
	e3 := size + size + size - 1
	stackSegments[0] = []int{s1, e1}
	stackSegments[1] = []int{s2, e2}
	stackSegments[2] = []int{s3, e3}
	stackPointers := make([]int, 3)
	stackPointers[0] = s1 - 1
	stackPointers[1] = s2 - 1
	stackPointers[2] = s3 - 1
	t.stackSegments = stackSegments
	t.stackPointers = stackPointers
	return &t
}

func (t *ThreeStackArray) substackIsEmpty(i int) bool {
	return t.stackPointers[i] == t.stackSegments[i][0]-1
}

func (t *ThreeStackArray) IsEmpty() bool {
	for i := range t.stackPointers {
		if !t.substackIsEmpty(i) {
			return false
		}
	}
	return true
}

func (t *ThreeStackArray) Push(stackIndex int, data int) error {
	if stackIndex < 0 || stackIndex > 3 {
		return fmt.Errorf("stackIndex %d is incorrect", stackIndex)
	}
	if t.stackPointers[stackIndex] < t.stackSegments[stackIndex][1] {
		t.stackPointers[stackIndex] += 1
		t.stack[t.stackPointers[stackIndex]] = data
		return nil
	} else {
		return fmt.Errorf("stack at %d is full", stackIndex)
	}
}

func (t *ThreeStackArray) Pop(stackIndex int) (int, error) {
	if stackIndex < 0 || stackIndex > 3 {
		return -1, fmt.Errorf("stackIndex %d is incorrect", stackIndex)
	}
	if t.substackIsEmpty(stackIndex) {
		return -1, fmt.Errorf("Can't Pop, stack %d is empty", stackIndex)
	}
	data := t.stack[t.stackPointers[stackIndex]]
	t.stackPointers[stackIndex] -= 1
	return data, nil
}

func (t *ThreeStackArray) Print() {
	b := bytes.Buffer{}
	for j, seg := range t.stackSegments {
		s := seg[0]
		e := seg[1]
		b.WriteString("|<")
		if !t.substackIsEmpty(j) {
			for i := s; i <= e; i++ {
				if t.stackPointers[j] == i {
					b.WriteString(fmt.Sprintf(" *%d ", t.stack[i]))
				} else {
					b.WriteString(fmt.Sprintf(" %d ", t.stack[i]))
				}
			}
		}
		b.WriteString(">|")
	}
	fmt.Println(b.String())
}

func main() {
	fmt.Println(problemTitle("3.1"))
	ts := NewThreeStackArray(10)
	ts.Print()
	ts.Push(0, 999)
	ts.Print()
	ts.Pop(0)
	ts.Print()
	ts.Push(0, 1111)
	ts.Print()
	ts.Push(1, 4235235)
	ts.Print()
	ts.Push(2, 910903)
	ts.Print()
	ts.Push(2, 34314)
	ts.Print()
	ts.Pop(2)
	ts.Print()
	ts.Push(2, 1112312)
	_, err := ts.Pop(2)
	if err != nil {
		fmt.Println(err)
	}
	_, err = ts.Pop(2)
	if err != nil {
		fmt.Println(err)
	}
	_, err = ts.Pop(2)
	if err != nil {
		fmt.Println(err)
	}
	ts.Print()
	fmt.Println(problemEndline(problemTitle("3.1")))

	fmt.Println(problemTitle("3.2"))
	fmt.Println(problemEndline(problemTitle("3.2")))
}
