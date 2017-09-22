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
}
