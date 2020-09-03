//Generated Test_isBalanced
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leetcode

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  9,
						Left: nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 2,
						Left:  &TreeNode{
							Val: 15,
							Left: &TreeNode{
								Val:  3,
								Left: nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val: 3,
								Left:  &TreeNode{Val: 4},
								Right: &TreeNode{Val: 4},
							},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 4,
								Left: nil,
								Right: nil,
							},
							Right: nil,
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val: 2,
						Left: nil,
						Right: &TreeNode{
							Val: 3,
							Left: nil,
							Right: &TreeNode{
								Val: 4,
								Left: nil,
								Right: nil,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
