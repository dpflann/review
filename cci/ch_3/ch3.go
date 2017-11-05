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

// ===== 3.2 =====
// create a stack with mim functionality
type MinStack struct {
	head *minStackNode
	tail *minStackNode // will be the top of the stack
}

type minStackNode struct {
	data     int
	min      int
	previous *minStackNode
	next     *minStackNode
}

func (ms *MinStack) Push(data int) {
	msn := &minStackNode{data: data}
	if ms.head == nil && ms.tail == nil {
		msn.min = data
		ms.head = msn
		ms.tail = msn
	} else {
		if data < ms.tail.min {
			msn.min = data
		} else {
			msn.min = ms.tail.min
		}
		msn.previous = ms.tail
		ms.tail.next = msn
		ms.tail = msn
	}
}

func (ms *MinStack) Pop() (int, error) {
	if ms.head == nil && ms.tail == nil {
		return -1, fmt.Errorf("stack is empty")
	}
	data := ms.tail.data
	if ms.tail == ms.head {
		ms.head = nil
	}
	ms.tail = ms.tail.previous
	return data, nil
}

func (ms *MinStack) Min() (int, error) {
	if ms.head == nil && ms.tail == nil {
		return -1, fmt.Errorf("stack is empty")
	}
	return ms.tail.min, nil
}

// ===== 3.3 =====
// SetOfStacks - will be comprised of multiple stacks
// Basic Stack

type Stack struct {
	capacity   int
	size       int
	topOfStack *Node
}

type Node struct {
	next     *Node
	previous *Node
	data     interface{}
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return fmt.Errorf("stack is empty")
	}
	data := s.topOfStack.data
	s.topOfStack = s.topOfStack.previous
	s.size -= 1
	return data
}

func (s *Stack) Peek() interface{} {
	if s.size == 0 {
		return fmt.Errorf("stack is empty")
	}
	return s.topOfStack.data
}

func (s *Stack) Push(data interface{}) error {
	if s.size == s.capacity {
		return fmt.Errorf("stack is full")
	}
	newNode := &Node{nil, nil, data}
	if s.topOfStack == nil {
		s.topOfStack = newNode
	} else {
		newNode.previous = s.topOfStack
		s.topOfStack = newNode
	}
	s.size += 1
	return nil
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsFull() bool {
	return s.size == s.capacity
}
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) String() string {
	base := ""
	if s.topOfStack == nil {
		return base
	}
	base = "top=" + fmt.Sprintf("%#v", s.topOfStack.data)
	node := s.topOfStack.previous
	for node != nil {
		base += "->" + fmt.Sprintf("%#v", node.data)
		node = node.previous
	}
	return base
}

type SetOfStacks struct {
	subStackCapacity int
	stacks           []Stack
	head             int
}

func (ss *SetOfStacks) Size() int {
	return len(ss.stacks)
}

func (ss *SetOfStacks) Push(data interface{}) error {
	if len(ss.stacks) == 0 {
		ss.head = 0
		ss.stacks = append(ss.stacks, Stack{ss.subStackCapacity, 0, nil})
	}
	if ss.stacks[ss.head].IsFull() {
		ss.head += 1
		ss.stacks = append(ss.stacks, Stack{ss.subStackCapacity, 0, nil})
	}
	return ss.stacks[ss.head].Push(data)
}

func (ss *SetOfStacks) Pop() interface{} {
	if len(ss.stacks) == 0 {
		return fmt.Errorf("Stack is empty")
	}
	for ss.stacks[ss.head].IsEmpty() && ss.head > 0 {
		ss.head -= 1
	}
	data := ss.stacks[ss.head].Pop()
	if ss.stacks[ss.head].IsEmpty() {
		ss.stacks = ss.stacks[:ss.head]
		ss.head -= 1
	}
	return data
}

func (ss *SetOfStacks) PopAt(i int) interface{} {
	if len(ss.stacks) == 0 {
		return fmt.Errorf("Stack is empty")
	}
	if i < 0 || i > len(ss.stacks) {
		return fmt.Errorf("index is out of range")
	}
	if i == ss.head {
		// Normal stack management
		return ss.Pop()
	} else {
		// May need to delete emptied stack
		data := ss.stacks[i].Pop()
		if ss.stacks[i].IsEmpty() {
			ss.stacks = append(ss.stacks[:i], ss.stacks[i+1:]...)
			// Shift the head, we have removed a stack.
			ss.head -= 1
		}
		return data
	}
}

// ===== 3.4 =====
// Towers of Hanoi
// 3 towers, N disks start in sorted order
// Rules:
// 1. only one disk can be moved at a time
// 2. a disk is slid off the top of one tower onto the next tower
// 3. a disk can only be placed on top of a larger disk
// Need to represent the game using stacks.
// Each tower is a stack with a position
// Start state: stack @ 0 has all disks
// End state: statck @ 2 (3rd tower) has all disks

func Hanoi(disk int, from, to, aux *Stack) {
	if disk <= 0 {
		return
	}
	Hanoi(disk-1, from, aux, to)
	to.Push(from.Pop())
	Hanoi(disk-1, aux, to, from)
}

// ===== 3.5 =====
// Implement a queue using two stacks
// LIFO to support FIFO
type Queue struct {
	enqueueStack Stack
	dequeueStack Stack
	capacity     int
}

func NewQueue(cap int) *Queue {
	eStack := Stack{cap, 0, nil}
	dStack := Stack{cap, 0, nil}
	return &Queue{eStack, dStack, cap}
}

func (q *Queue) Enqueue(data interface{}) error {
	if q.IsFull() {
		return fmt.Errorf("Queue is full")
	}
	err := q.enqueueStack.Push(data)
	if err != nil {
		return fmt.Errorf("Unable to enqueue data")
	}
	return nil
}

func (q *Queue) Dequeue() interface{} {
	if q.dequeueStack.IsEmpty() && q.enqueueStack.IsEmpty() {
		return nil
	}
	if q.dequeueStack.IsEmpty() && !q.enqueueStack.IsEmpty() {
		// move all enqueued items to the dequeueStack
		for !q.enqueueStack.IsEmpty() {
			q.dequeueStack.Push(q.enqueueStack.Pop())
		}
	}
	return q.dequeueStack.Pop()
}

func (q *Queue) IsFull() bool {
	return q.enqueueStack.IsFull() && q.dequeueStack.IsFull()
}

// ===== 3.3 =====
// Sort Stack

///\\\///\\\///[[ Begin Testing ]]\\\///\\\//\\\
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
	ms := MinStack{}
	ms.Push(5)
	min, _ := ms.Min()
	fmt.Println(min == 5)
	ms.Push(10)
	min, _ = ms.Min()
	fmt.Println(min == 5)
	ms.Push(1)
	min, _ = ms.Min()
	fmt.Println(min == 1)
	min, _ = ms.Min()
	ms.Pop()
	min, _ = ms.Min()
	ms.Pop()
	min, _ = ms.Min()
	fmt.Println(min == 5)
	ms.Pop()
	_, err = ms.Pop()
	if err != nil {
		fmt.Println(err.Error() == "stack is empty")
	}
	min, _ = ms.Min()
	fmt.Println(min == -1)
	fmt.Println(problemEndline(problemTitle("3.2")))

	fmt.Println(problemTitle("3.3"))
	// test the simple stack
	s := Stack{5, 0, nil}
	fmt.Println(s.Pop())
	fmt.Println(s.Push(50))
	fmt.Println(s.Size() == 1)
	fmt.Println(s.Pop())
	// Fill Up Stack
	fmt.Println(s.Push(12))
	fmt.Println(s.Push(1))
	fmt.Println(s.Push(654))
	fmt.Println(s.Push(12300))
	fmt.Println(s.Push("hello"))
	// At capacity - expect an error
	fmt.Println(s.Push(53))

	fmt.Println(problemTitle("SetOfStacks"))
	ss := SetOfStacks{2, nil, 0}
	fmt.Println("Push(10)", ss.Push(10))
	fmt.Println("Size == 1", ss.Size() == 1)
	fmt.Println("Pop() --> 10", ss.Pop() == 10)
	fmt.Println("Push(11)", ss.Push(11))
	fmt.Println("Push(12)", ss.Push(12))
	fmt.Println("Push(13)", ss.Push(13))
	fmt.Println("Size == 2", ss.Size() == 2)
	fmt.Println("Pop() --> 13", ss.Pop() == 13)
	fmt.Println("Push(14)", ss.Push(14))
	fmt.Println("Push(15)", ss.Push(15))
	fmt.Println("Push(16)", ss.Push(16))
	fmt.Println("Size == 3", ss.Size() == 3)
	fmt.Println("Pop() -- > 16", ss.Pop() == 16)
	fmt.Println("Pop() --> 15", ss.Pop() == 15)
	fmt.Println("Pop() --> 14", ss.Pop() == 14)
	fmt.Println("Pop() --> 12", ss.Pop() == 12)
	fmt.Println("Pop() --> 11", ss.Pop() == 11)
	fmt.Println("Pop() --> error", ss.Pop())
	fmt.Println(ss.Size() == 0)
	for i := range []int{0, 1, 2, 3, 4} {
		ss.Push(i)
	}
	fmt.Println(ss.Size() == 3)
	fmt.Println(ss.PopAt(1) == 3)
	fmt.Println(ss.PopAt(1) == 2)
	fmt.Println(ss.Size() == 2)
	fmt.Println(ss.Pop() == 4)
	fmt.Println(ss.Size() == 1)
	fmt.Println(ss.Pop() == 1)
	fmt.Println(ss.Pop() == 0)
	fmt.Println(ss.Size() == 0)
	fmt.Println(problemEndline(problemTitle("3.3")))

	fmt.Println(problemTitle("3.4"))
	firstPeg, secondPeg, thirdPeg := &Stack{5, 0, nil}, &Stack{5, 0, nil}, &Stack{5, 0, nil}
	for _, i := range []int{5, 4, 3, 2, 1} {
		firstPeg.Push(i)
	}
	Hanoi(5, firstPeg, thirdPeg, secondPeg)
	fmt.Println("First peg is empty:", firstPeg.IsEmpty() == true)
	fmt.Println("second peg is empty:", secondPeg.IsEmpty() == true)
	fmt.Println("third peg is full:", thirdPeg.IsFull() == true)
	fmt.Println("third peg: ", thirdPeg)
	for _, i := range []int{1, 2, 3, 4, 5} {
		fmt.Println(thirdPeg.Pop() == i)
	}
	fmt.Println(problemEndline(problemTitle("3.4")))

	fmt.Println(problemTitle("3.5"))
	q := NewQueue(5)
	for i := range []int{0, 1, 2, 3, 4} {
		fmt.Println(q.Enqueue(i))
	}
	fmt.Println(q.Enqueue(5))
	fmt.Println(q.Dequeue())
	fmt.Println(q.Enqueue(5))
	for _, v := range []int{1, 2, 3, 4, 5} {
		fmt.Println(q.Dequeue() == v)
	}
	fmt.Println(q.Dequeue())
	fmt.Println(problemEndline(problemTitle("3.5")))

	fmt.Println(problemTitle("3.6"))
	fmt.Println(problemEndline(problemTitle("3.6")))
}
