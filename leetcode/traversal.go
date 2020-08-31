package leetcode

import(
	"container/list"
	"fmt"
)

func traversal(node *TreeNode) {
	stack := list.New()
	// 前序遍历
	if node != nil {
		stack.PushBack(node)
	}
	cur := node
	for stack.Len() != 0 {
		r := stack.Back()
		cur = r.Value.(*TreeNode)
		stack.Remove(r)
		fmt.Println(cur.Right)
		fmt.Println(cur.Left)
		if cur.Right != nil || cur.Left != nil {
			if cur.Right != nil { stack.PushBack(cur.Right) }
			if cur.Left != nil { stack.PushBack(cur.Left) }
		} else {
		}
	}
}