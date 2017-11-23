package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
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

// ===== 4.2 =====
// This current Graph concept has little
// appreciation for edges, edge weight, and explicit concept
// of vertex.

type GNode struct {
	Adjacents []*GNode
	Data      interface{}
	Visited   bool
}

func (n *GNode) AddAdjacent(adjacentNode *GNode) {
	if adjacentNode == nil {
		return
	}
	if n.Adjacents == nil {
		n.Adjacents = []*GNode{adjacentNode}
	} else {
		for _, an := range n.Adjacents {
			if an == adjacentNode {
				return
			}
		}
		n.Adjacents = append(n.Adjacents, adjacentNode)
	}
}

func (n *GNode) RemoveAdjacent(gNode *GNode) {
	if gNode == nil {
		return
	}
	if len(n.Adjacents) > 0 {
		for i := range n.Adjacents {
			if n.Adjacents[i] == gNode {
				n.Adjacents = append(n.Adjacents[0:i], n.Adjacents[i+1:]...)
			}
		}
	}
}

func (n *GNode) IsAdjacentTo(gNode *GNode) bool {
	if gNode == nil || n.Adjacents == nil {
		return false
	}
	for _, node := range n.Adjacents {
		if gNode == node {
			return true
		}
	}
	return false
}

func (n *GNode) String() string {
	s := fmt.Sprintf("Node: %v\n", n.Data)
	if n.Adjacents == nil {
		return s
	}
	for _, adj := range n.Adjacents {
		s += fmt.Sprintf("\t-> %v\n", adj.Data)
	}
	return s
}

type Graph struct {
	Nodes []*GNode
}

func (g *Graph) HasPath(start, end *GNode) bool {
	// Perform BFS and check adjacenct nodes of current
	// node for existence of end.
	if start == nil || end == nil {
		return false
	}
	nodesQ := []*GNode{}
	start.Visited = true
	nodesQ = append(nodesQ, start)
	for len(nodesQ) > 0 {
		currentNode := nodesQ[0]
		nodesQ = nodesQ[1:]
		for _, n := range currentNode.Adjacents {
			if n == end {
				return true
			}
			if !n.Visited {
				n.Visited = true
				nodesQ = append(nodesQ, n)
			}
		}
	}
	return false
}

func (g *Graph) String() string {
	s := "Graph:\n"
	if g.Nodes == nil {
		return s
	}
	for _, gn := range g.Nodes {
		s += gn.String()
	}
	return s
}

// ===== 4.3 =====
// Construct minimal height binary search tree from array in ascending sorted order
// [1, 2, 3, 4, 5, 6, 7]
//       4
//    /     \
//   2       6
//  / \     / \
// 1   3   5   7

// MinimalHeightBST converts an array in ascending order into
// a binary search tree with minimal height.
// A *BNode is returned which is the root of the BST.
func MinimalHeightBST(sortedArray []int) *BNode {
	if len(sortedArray) == 0 {
		return nil
	}
	if len(sortedArray) == 1 {
		return &BNode{
			Data:  sortedArray[0],
			Left:  nil,
			Right: nil,
		}
	}

	index := len(sortedArray) / 2
	root := &BNode{}
	root.Data = sortedArray[index]

	root.Left = MinimalHeightBST(sortedArray[0:index])
	root.Right = MinimalHeightBST(sortedArray[index+1:])

	return root
}

//\\//\\//\\ Main - Testing Solutions //\\//\\//\\
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

	fmt.Println(problemTitle("4.2"))
	// Make nodes and graph to test `HasPath`
	graph := Graph{}
	for i := 0; i < 10; i++ {
		n := GNode{
			Adjacents: nil,
			Data:      i,
			Visited:   false,
		}
		graph.Nodes = append(graph.Nodes, &n)
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	start, end := &GNode{}, &GNode{}
	for i := 10; i < 20; i++ {
		// directed graph
		// pick nodes from the graph and create connections
		i := r1.Int31n(10)
		j := r1.Int31n(10)
		for i == j {
			j = r1.Int31n(10)
		}
		node1 := graph.Nodes[i]
		node2 := graph.Nodes[j]
		// node1 --> node2
		node1.AddAdjacent(node2)
	}
	// Ensure valid start, end
	start = graph.Nodes[0]
	end = graph.Nodes[9]
	start.AddAdjacent(end)
	fmt.Printf("Connected start (%v) to end (%v)\n", start.Data, end.Data)
	fmt.Println(graph)
	fmt.Println("graph.HasPath(start, end)?")
	fmt.Println(graph.HasPath(start, end) == true)
	loneNode := &GNode{
		Adjacents: nil,
		Data:      111111,
		Visited:   false,
	}
	fmt.Println(graph.HasPath(start, loneNode) == false)
	fmt.Println(problemEndline(problemTitle("4.2")))

	fmt.Println(problemTitle("4.3"))
	sortedArrayEvenLength := []int{0, 1, 2, 3, 4, 5, 6, 7}
	expectedHeight := 4
	eBst := MinimalHeightBST(sortedArrayEvenLength)
	inOrder := ""
	InOrder(eBst, &inOrder)
	fmt.Println(inOrder)
	fmt.Println("eBST has expected height:", Height(eBst) == expectedHeight)

	sortedArrayOddLength := []int{0, 1, 2, 3, 4, 5, 6}
	expectedHeight = 3
	oBst := MinimalHeightBST(sortedArrayOddLength)
	inOrder = ""
	InOrder(oBst, &inOrder)
	fmt.Println(inOrder)
	fmt.Println("oBST has expected height:", Height(oBst) == expectedHeight)

	// Edge Cases
	edgeCaseBST := MinimalHeightBST([]int{})
	fmt.Println(edgeCaseBST == nil)
	edgeCaseBST = MinimalHeightBST([]int{0})
	fmt.Println(Height(edgeCaseBST) == 1)

	fmt.Println(problemEndline(problemTitle("4.3")))
}
