package main

import "fmt"

type Node struct {
	left, right *Node
	value       int
}

type Tree struct {
	root *Node
}

func (t *Tree) Add(value int) {
	if t.root == nil {
		t.root = &Node{
			value: value,
		}
	} else {
		cur := t.root
		var prev *Node
		for cur != nil {
			prev = cur
			if cur.value < value {
				cur = cur.left
			} else if cur.value > value {
				cur = cur.right
			} else {
				return
			}
		}

		if prev.value < value {
			prev.left = &Node{
				value: value,
			}
		} else {
			prev.right = &Node{
				value: value,
			}
		}
	}
}

func (t *Tree) Depth() int {
	return t.depth(t.root, 0)
}

func (t *Tree) depth(node *Node, cur_depth int) int {
	if node != nil {
		cur_depth++
		lcount := t.depth(node.left, cur_depth)
		rcount := t.depth(node.right, cur_depth)

		if lcount > rcount {
			return lcount
		} else {
			return rcount
		}
	}

	return cur_depth
}

func main() {
	var x int
	tree := Tree{}

	fmt.Scan(&x)

	for x != 0 {
		tree.Add(x)
		fmt.Scan(&x)
	}

	fmt.Println(tree.Depth())

}
