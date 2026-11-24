package main

import "fmt"

// TreeNode is a binary search tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// LCA returns the lowest common ancestor of values a and b in a BST (both
// assumed present): descend while both values are on the same side; the
// first node between them is the answer.
func LCA(root *TreeNode, a, b int) *TreeNode {
	for root != nil {
		switch {
		case a < root.Val && b < root.Val:
			root = root.Left
		case a > root.Val && b > root.Val:
			root = root.Right
		default:
			return root // split point, or equal to one of the values
		}
	}
	return nil
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(6, tn(2, tn(0, nil, nil), tn(4, tn(3, nil, nil), tn(5, nil, nil))), tn(8, tn(7, nil, nil), tn(9, nil, nil)))
	fmt.Println(LCA(root, 2, 8).Val, LCA(root, 3, 5).Val)
}
