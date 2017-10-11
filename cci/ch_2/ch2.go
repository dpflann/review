package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Next *Node
	Data int
}

type LinkedList struct {
	Head *Node
}

func (LL *LinkedList) AddNode(newNode *Node) {
	if LL == nil {
		// initialize
		LL = &LinkedList{newNode}
	} else if LL.Head == nil {
		LL.Head = newNode
	} else {
		currentNode := LL.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}
}

func (Llist1 *LinkedList) Equals(Llist2 *LinkedList) bool {
	if Llist1 == nil && Llist2 == nil {
		return true
	}
	if Llist1 == nil || Llist2 == nil {
		return false
	}
	currentNode := Llist1.Head
	currentNoNode := Llist2.Head
	for currentNode != nil && currentNoNode != nil {
		if currentNode.Data != currentNoNode.Data {
			return false
		}
		currentNode = currentNode.Next
		currentNoNode = currentNoNode.Next
	}
	return true
}

func (LL *LinkedList) String() string {
	currentNode := LL.Head
	strLL := ""
	for currentNode != nil {
		strLL += strconv.Itoa(currentNode.Data) + "->"
		currentNode = currentNode.Next
	}
	return strLL
}

func RemoveDuplicates1(LL *LinkedList) *LinkedList {
	if LL == nil || LL.Head == nil {
		return nil
	}
	uniqueData := make(map[int]bool)
	currentNode := LL.Head
	previousNode := LL.Head
	for currentNode != nil {
		_, seen := uniqueData[currentNode.Data]
		if seen {
			previousNode.Next = currentNode.Next
			currentNode = previousNode.Next
		} else {
			uniqueData[currentNode.Data] = true
			previousNode = currentNode
			currentNode = currentNode.Next
		}
	}
	return LL
}

func RemoveDuplicates2(LL *LinkedList) *LinkedList {
	if LL == nil || LL.Head == nil {
		return nil
	}
	currentNode := LL.Head
	for currentNode != nil {
		previousNode := currentNode
		nextNode := currentNode.Next
		for nextNode != nil {
			if nextNode.Data == currentNode.Data {
				previousNode.Next = nextNode.Next
				previousNode = previousNode.Next
				if previousNode != nil {
					nextNode = nextNode.Next.Next
				} else {
					nextNode = nil
				}
			} else {
				previousNode = nextNode
				nextNode = nextNode.Next
			}
		}
		currentNode = currentNode.Next
	}
	return LL
}

func FindKthToLast(LL *LinkedList, k int) *Node {
	length := 0
	for currentNode := LL.Head; currentNode != nil; currentNode = currentNode.Next {
		length += 1
	}
	index := length - k - 1
	i := 0
	currentNode := LL.Head
	for i < index {
		currentNode = currentNode.Next
		i += 1
	}
	return currentNode
}

func DeleteNode(n *Node) {
	if n == nil {
		return
	}
	if n.Next != nil {
		tmpNode := n.Next
		n.Next = tmpNode.Next
		n.Data = tmpNode.Data
	}
}

func Partition(p int, LL *LinkedList) *LinkedList {
	if LL == nil || LL.Head == nil {
		return nil
	}
	lesserList, equalAndGreaterList := &LinkedList{}, &LinkedList{}
	currentNode := LL.Head
	for currentNode != nil {
		newNode := currentNode.Next
		if currentNode.Data < p {
			currentNode.Next = nil
			lesserList.AddNode(currentNode)
		} else {
			equalAndGreaterList.AddNode(currentNode)
		}
		currentNode.Next = nil
		currentNode = newNode
	}
	lesserList.AddNode(equalAndGreaterList.Head)
	return lesserList
}

func AddNumbersAsLinkedLists(LL1, LL2 *LinkedList) *LinkedList {
	sum := NumberFromLinkedList(LL1) + NumberFromLinkedList(LL2)
	return LinkedListFromNumber(sum)
}

// 7->1->6 = 617, digits are stored in reverse order
func NumberFromLinkedList(LL *LinkedList) int {
	if LL.Head == nil {
		return -1
	}
	total, power := LL.Head.Data, 10
	currentNode := LL.Head.Next
	for currentNode != nil {
		total += currentNode.Data * power
		power = power * 10
		currentNode = currentNode.Next
	}
	return total
}

func LinkedListFromNumber(n int) *LinkedList {
	LL := &LinkedList{}
	for n > 0 {
		digit := n % 10
		n = n / 10
		LL.AddNode(&Node{nil, digit})
	}
	return LL
}

func AddNumbersAsLinkedListsReversed(LL1, LL2 *LinkedList) *LinkedList {
	sum := NumberFromLinkedListReversed(LL1) + NumberFromLinkedListReversed(LL2)
	return LinkedListFromNumberReversed(sum)
}

// 7->1->6 = 617, digits are stored in reverse order
func NumberFromLinkedListReversed(LL *LinkedList) int {
	if LL.Head == nil {
		return -1
	}
	total := 0
	currentNode := LL.Head
	for currentNode != nil {
		total = total*10 + currentNode.Data
		currentNode = currentNode.Next
	}
	return total
}

func LinkedListFromNumberReversed(n int) *LinkedList {
	LL := &LinkedList{}
	var currentNode, previousNode *Node
	for n > 0 {
		digit := n % 10
		n = n / 10
		currentNode = &Node{nil, digit}
		if previousNode == nil {
			previousNode = currentNode
		} else {
			currentNode.Next = previousNode
			previousNode = currentNode
		}
	}
	return LL
}

func FindCycleStart(LL *LinkedList) *Node {
	// A -> -B -> C -> D -> E -> C
	seenNodes := map[*Node]bool{}
	currentNode := LL.Head
	for currentNode != nil {
		if _, ok := seenNodes[currentNode]; ok {
			return currentNode
		} else {
			seenNodes[currentNode] = true
		}
		currentNode = currentNode.Next
	}
	return nil
}

func FindCycleStartV2(LL *LinkedList) *Node {
	// A -> -B -> C -> D -> E -> C
	seenNodes := map[*Node]bool{}
	currentNode := LL.Head
	for currentNode != nil {
		if _, ok := seenNodes[currentNode]; ok {
			return currentNode
		} else {
			seenNodes[currentNode] = true
		}
		currentNode = currentNode.Next
	}
	return nil
}

func (LL *LinkedList) IsPalindrome() bool {
	nodes := []int{}
	currentNode := LL.Head
	for currentNode != nil {
		nodes = append(nodes, currentNode.Data)
		currentNode = currentNode.Next
	}
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		if nodes[i] != nodes[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("====== 2.1 ======")
	list1, dupeList1, expectedList1 := &LinkedList{}, &LinkedList{}, &LinkedList{}
	for i := 0; i < 10; i++ {
		n1 := &Node{nil, i}
		n2 := &Node{nil, i}
		n3 := &Node{nil, i}
		n4 := &Node{nil, i}
		list1.AddNode(n1)
		expectedList1.AddNode(n2)
		dupeList1.AddNode(n3)
		dupeList1.AddNode(n4)
	}

	fmt.Println("RemoveDuplicates1(...)")
	fmt.Println(RemoveDuplicates1(list1).Equals(expectedList1) == true)
	fmt.Println(RemoveDuplicates1(dupeList1).Equals(expectedList1) == true)

	list1, dupeList1, expectedList1 = &LinkedList{}, &LinkedList{}, &LinkedList{}
	for i := 0; i < 10; i++ {
		n1 := &Node{nil, i}
		n2 := &Node{nil, i}
		n3 := &Node{nil, i}
		n4 := &Node{nil, i}
		list1.AddNode(n1)
		expectedList1.AddNode(n2)
		dupeList1.AddNode(n3)
		dupeList1.AddNode(n4)
	}
	fmt.Println("RemoveDuplicates2(...)")
	fmt.Println(RemoveDuplicates2(list1).Equals(expectedList1) == true)
	fmt.Println(RemoveDuplicates2(dupeList1).Equals(expectedList1) == true)
	fmt.Println("=================")

	fmt.Println("====== 2.2 ======")
	fmt.Println("FindKthToLast(..., ...)")
	fmt.Println(list1)
	fmt.Println(FindKthToLast(list1, 0).Data == 9)
	fmt.Println(FindKthToLast(list1, 9).Data == 0)
	fmt.Println(FindKthToLast(list1, 8).Data == 1)
	fmt.Println("=================")

	fmt.Println("====== 2.3 ======")
	list2, expectedList2 := &LinkedList{}, &LinkedList{}
	for i := 0; i < 10; i++ {
		n1 := &Node{nil, i}
		n2 := &Node{nil, i}
		list2.AddNode(n1)
		expectedList2.AddNode(n2)
	}
	n1 := &Node{nil, 11}
	list2.AddNode(n1)
	for i := 12; i < 15; i++ {
		n1 := &Node{nil, i}
		n2 := &Node{nil, i}
		list2.AddNode(n1)
		expectedList2.AddNode(n2)
	}

	fmt.Println("DeleteNode(...)")
	DeleteNode(n1)
	fmt.Println(list2.Equals(expectedList2) == true)
	fmt.Println("=================")

	fmt.Println("====== 2.4 ======")
	list3, expectedList3 := &LinkedList{}, &LinkedList{}
	n1 = &Node{nil, 11}
	for i := 0; i < 10; i++ {
		n1 := &Node{nil, i}
		n2 := &Node{nil, i}
		list3.AddNode(n1)
		expectedList3.AddNode(n2)
	}
	// TODO: dpf.2017.09.20 -- instantiate a Node pointer in AddNode to simplify code
	expectedList3.AddNode(&Node{nil, 9})
	for i := 12; i < 15; i++ {
		n1 := &Node{nil, i}
		n2 := &Node{nil, i}
		list3.AddNode(n1)
		expectedList3.AddNode(n2)
	}
	list3.AddNode(&Node{nil, 9})
	fmt.Println("Partition(...)")
	partitionedList := Partition(10, list3)
	fmt.Println(partitionedList.Equals(expectedList3) == true)
	fmt.Println("=================")

	n, m, expectedSum, expectedSumReversed := &LinkedList{}, &LinkedList{}, &LinkedList{}, &LinkedList{}

	// 7->1->6 = 617
	n.AddNode(&Node{nil, 7})
	n.AddNode(&Node{nil, 1})
	n.AddNode(&Node{nil, 6})

	// 5->4->3 = 345
	m.AddNode(&Node{nil, 5})
	m.AddNode(&Node{nil, 4})
	m.AddNode(&Node{nil, 3})

	// 617 + 345 =  962 == 2->6->9
	expectedSum.AddNode(&Node{nil, 2})
	expectedSum.AddNode(&Node{nil, 6})
	expectedSum.AddNode(&Node{nil, 9})

	fmt.Println("====== 2.5 ======")
	fmt.Println("AddNumbersAsLinkedLists(..., ...)")
	fmt.Println(NumberFromLinkedList(n) == 617)
	fmt.Println(NumberFromLinkedList(m) == 345)
	fmt.Println(NumberFromLinkedList(expectedSum) == 962)
	fmt.Println(LinkedListFromNumber(617).Equals(n))
	fmt.Println(LinkedListFromNumber(345).Equals(m))
	fmt.Println(LinkedListFromNumber(962).Equals(expectedSum))
	fmt.Println(AddNumbersAsLinkedLists(n, m).Equals(expectedSum))

	// "reversed order"
	expectedSumReversed.AddNode(&Node{nil, 1})
	expectedSumReversed.AddNode(&Node{nil, 2})
	expectedSumReversed.AddNode(&Node{nil, 5})
	expectedSumReversed.AddNode(&Node{nil, 9})

	fmt.Println(NumberFromLinkedListReversed(n) == 716)
	fmt.Println(NumberFromLinkedListReversed(m) == 543)

	fmt.Println(NumberFromLinkedListReversed(expectedSumReversed) == 1259)
	fmt.Println(LinkedListFromNumberReversed(716).Equals(n))
	fmt.Println(LinkedListFromNumberReversed(543).Equals(m))
	fmt.Println(LinkedListFromNumberReversed(1259).Equals(expectedSumReversed))
	fmt.Println(AddNumbersAsLinkedListsReversed(n, m).Equals(expectedSumReversed))
	fmt.Println("=================")

	fmt.Println("====== 2.6 ======")
	fmt.Println("FindCycleStart(...)")
	cycle, noCycle := &LinkedList{}, &LinkedList{}
	repeatedNode := &Node{nil, 7}

	cycle.AddNode(&Node{nil, 10})
	cycle.AddNode(&Node{nil, 1})
	cycle.AddNode(repeatedNode)
	cycle.AddNode(&Node{nil, 6})
	cycle.AddNode(&Node{nil, 4})
	cycle.AddNode(repeatedNode)

	noCycle.AddNode(&Node{nil, 10})
	noCycle.AddNode(&Node{nil, 1})
	noCycle.AddNode(&Node{nil, 6})
	noCycle.AddNode(&Node{nil, 4})

	fmt.Println(FindCycleStart(cycle) != nil)
	fmt.Println(FindCycleStart(cycle).Data == repeatedNode.Data)
	fmt.Println(FindCycleStart(cycle) == repeatedNode)
	fmt.Println(FindCycleStart(noCycle) == nil)
	fmt.Println("=================")

	fmt.Println("====== 2.7 ======")
	fmt.Println("IsPalindrome()")
	palindrome, noPalindrome := &LinkedList{}, &LinkedList{}

	palindrome.AddNode(&Node{nil, 10})
	palindrome.AddNode(&Node{nil, 1})
	palindrome.AddNode(&Node{nil, 6})
	palindrome.AddNode(&Node{nil, 4})
	palindrome.AddNode(&Node{nil, 6})
	palindrome.AddNode(&Node{nil, 1})
	palindrome.AddNode(&Node{nil, 10})

	noPalindrome.AddNode(&Node{nil, 10})
	noPalindrome.AddNode(&Node{nil, 6})
	noPalindrome.AddNode(&Node{nil, 4})
	noPalindrome.AddNode(&Node{nil, 6})

	fmt.Println(palindrome.IsPalindrome() == true)
	fmt.Println(noPalindrome.IsPalindrome() == false)
	fmt.Println("=================")
}
