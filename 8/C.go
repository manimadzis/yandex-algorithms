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

func (t *Tree) Max(i int) int {
	var maxx int
	t.max(t.root, &i, &maxx)
	return maxx
}

func (t *Tree) max(node *Node, i *int, maxx *int) {
	if node == nil {
		return
	}
	t.max(node.right, i, maxx)

	if *i > 0 {
		(*i)--
		*maxx = node.value
	}

	t.max(node.left, i, maxx)
}

func main() {
	var x int
	tree := Tree{}

	fmt.Scan(&x)

	for x != 0 {
		tree.Add(x)
		fmt.Scan(&x)
	}

	fmt.Println(tree.Max(2))
}
