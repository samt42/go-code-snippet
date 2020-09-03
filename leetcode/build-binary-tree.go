package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var rootIndex int
	for i:=0; i<len(inorder); i++ {
		if inorder[i] == preorder[0] {
			rootIndex = i
			break
		}
	}
	var root *TreeNode = &TreeNode{Val:inorder[rootIndex]}
	fmt.Println(preorder)
	fmt.Println(inorder)
	fmt.Println(rootIndex)
	if len(preorder) > 1 {
		root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
		root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	}
	return root
}

