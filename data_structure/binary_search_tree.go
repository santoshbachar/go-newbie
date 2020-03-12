package main

import "fmt"

type Tree struct {
	node *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (t *Tree) insert(val int) {
	if t.node == nil {
		t.node = &Node{data: val, left: nil, right: nil}
	} else {
		t.node.insert(val)
	}
}

func (n *Node) insert(val int) {
	if n == nil {
		return
	}
	if val <= n.data {
		if n.left == nil {
			n.left = &Node{data: val, left: nil, right: nil}
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: val, left: nil, right: nil}
		} else {
			n.right.insert(val)
		}
	}
}

func printInOrder(n *Node) {
	if n == nil {
		return
	}
	printInOrder(n.left)
	fmt.Println(n.data)
	printInOrder(n.right)
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.data)
	printPreOrder(n.left)
	printPreOrder(n.right)
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	}
	printPostOrder(n.left)
	printPostOrder(n.right)
	fmt.Println(n.data)
}

func main() {
	t := Tree{}
	t.insert(100)
	t.insert(89)
	t.insert(78)
	t.insert(99)
	t.insert(34)

	fmt.Println("Inorder")
	printInOrder(t.node)
	fmt.Println("Preorder")
	printPreOrder(t.node)
	fmt.Println("Postorder")
	printPostOrder(t.node)
}
