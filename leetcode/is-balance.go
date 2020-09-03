/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leetcode

import (
	"math"
)

func isBalanced(root *TreeNode) bool {
	return depth(root) >=0
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	var leftDepth float64 = float64( depth(node.Left))
	var rightDepth float64 = float64(depth(node.Right))
	if leftDepth >= 0 && rightDepth >= 0 && math.Abs(leftDepth - rightDepth) <= 1 {
		return int(math.Max(leftDepth, rightDepth)) + 1
	} else {
		return -1
	}
}


