/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	cur := root
	depth := 0
	for cur.Right != nil && cur.Left != nil {
		if cur.Right != nil {
			cur = cur.Right
			depth++
		} else if cur.Left != nil {
			cur = cur.Left
			depth++
		}
	}
	return depth
}


