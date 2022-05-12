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
			if cur.value < value {
				cur = cur.left
			} else if cur.value > value {
				cur = cur.right
			} else {
				return -1
			}
			depth++
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

	return depth
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
	lengths := []int{}

	fmt.Scan(&x)

	for x != 0 {
		depth := tree.Add(x)
		if depth > 0 {
			lengths = append(lengths, depth)
		}
		fmt.Scan(&x)
	}

	for _, v := range lengths {
		fmt.Printf("%d ", v)
	}

}
