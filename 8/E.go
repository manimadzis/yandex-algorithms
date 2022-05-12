package main

import "fmt"

type Node struct {
	left, right *Node
	value       int
}

type Tree struct {
	root *Node
}

func (t *Tree) Add(value int) int {
	depth := 1

	if t.root == nil {
		t.root = &Node{
			value: value,
		}
	} else {
		cur := t.root
		var prev *Node
		for cur != nil {
			prev = cur
			if cur.value > value {
				cur = cur.left
			} else if cur.value < value {
				cur = cur.right
			} else {
				return -1
			}
			depth++
		}

		if prev.value > value {
			prev.left = &Node{
				value: value,
			}
		} else {
			prev.right = &Node{
				value: value,
			}
		}
	}

	return depth
}

func (t *Tree) PrintLeaf() {
	t.printLeaf(t.root)
}

func (t *Tree) printLeaf(node *Node) {
	if node == nil {
		return
	}

	t.printLeaf(node.left)

	if node.left == nil && node.right == nil {
		fmt.Println(node.value)
	}

	t.printLeaf(node.right)
}

func main() {
	var x int
	tree := Tree{}

	fmt.Scan(&x)

	for x != 0 {
		tree.Add(x)
		fmt.Scan(&x)
	}

	tree.PrintLeaf()
}
