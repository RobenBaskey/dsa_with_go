package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	data := CreateTree()
	result := preorderTraversal(data)
	fmt.Println("My result is = ", result)
}

func preorderTraversal(root *Node) []int {
	var result []int
	var helper func(*Node)
	helper = func(node *Node) {
		if node == nil {
			return
		}
		result = append(result, node.data)
		helper(node.left)
		helper(node.right)
	}
	helper(root)
	return result
}

func CreateTree() *Node {
	two := CreateNewNode(2)
	seven := CreateNewNode(7)
	nine := CreateNewNode(9)
	AddLeftChild(two, seven)
	AddRightChild(two, nine)

	one := CreateNewNode(1)
	six := CreateNewNode(6)
	AddLeftChild(seven, one)
	AddRightChild(seven, six)

	eight := CreateNewNode(8)
	AddRightChild(nine, eight)

	three := CreateNewNode(3)
	four := CreateNewNode(4)
	AddLeftChild(eight, three)
	AddRightChild(eight, four)

	return two
}

func CreateNewNode(data int) *Node {
	var new_node Node

	new_node.data = data
	new_node.left = nil
	new_node.right = nil

	return &new_node
}

func AddLeftChild(node *Node, child *Node) {
	node.left = child
}

func AddRightChild(node *Node, child *Node) {
	node.right = child
}
