package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
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

// Binary Tree Structs
type BNode struct {
	Data  int
	Left  *BNode
	Right *BNode
}

func NewBNode(data int) *BNode {
	return &BNode{data, nil, nil}
}

func InOrder(root *BNode, str *string) {
	if root == nil {
		return
	}
	InOrder(root.Left, str)
	*str += strconv.Itoa(root.Data) + " -> "
	InOrder(root.Right, str)
}

func PreOrder(root *BNode, str *string) {
	if root == nil {
		return
	}
	*str += strconv.Itoa(root.Data) + " -> "
	PreOrder(root.Left, str)
	PreOrder(root.Right, str)
}

func PostOrder(root *BNode, str *string) {
	if root == nil {
		return
	}
	PostOrder(root.Left, str)
	PostOrder(root.Right, str)
	*str += strconv.Itoa(root.Data) + " -> "
}

// ===== 4.1 =====
func Height(bnode *BNode) int {
	if bnode == nil {
		return 0
	}
	return int(math.Max(float64(Height(bnode.Left)), float64(Height(bnode.Right)))) + 1
}

func IsBalanced(bnode *BNode) bool {
	if bnode == nil {
		return true
	}
	heightDiff := int(math.Abs(float64(Height(bnode.Left) - Height(bnode.Right))))
	if heightDiff > 1 {
		return false
	} else {
		return IsBalanced(bnode.Left) && IsBalanced(bnode.Right)
	}
}

func main() {
	fmt.Println(problemTitle("4.1"))
	fmt.Println(IsBalanced(nil) == true)
	// make a balanced tree
	root := NewBNode(100)
	root.Left = NewBNode(1)
	root.Right = NewBNode(2)
	root.Left.Left = NewBNode(3)
	root.Left.Right = NewBNode(4)
	root.Right.Left = NewBNode(5)
	root.Right.Right = NewBNode(5)
	root.Left.Left.Left = NewBNode(6)
	root.Left.Left.Right = NewBNode(7)
	root.Left.Right.Left = NewBNode(8)
	root.Left.Right.Right = NewBNode(9)
	root.Right.Left.Left = NewBNode(10)
	root.Right.Left.Right = NewBNode(11)
	root.Right.Right.Left = NewBNode(12)
	root.Right.Right.Right = NewBNode(13)

	fmt.Println("In Order")
	result := ""
	InOrder(root, &result)
	fmt.Println(result)

	fmt.Println("Pre Order")
	result = ""
	PreOrder(root, &result)
	fmt.Println(result)

	fmt.Println("Post Order")
	result = ""
	PostOrder(root, &result)
	fmt.Println(result)

	fmt.Println(IsBalanced(root) == true)
	// make an unbalanced tree
	root = NewBNode(100)
	root.Left = NewBNode(1)
	root.Left.Left = NewBNode(3)
	root.Left.Right = NewBNode(4)
	root.Left.Left.Left = NewBNode(6)
	root.Left.Left.Right = NewBNode(7)
	root.Left.Right.Left = NewBNode(8)
	root.Left.Right.Right = NewBNode(9)
	fmt.Println(IsBalanced(root) == false)
	fmt.Println(problemEndline(problemTitle("4.1")))
}
