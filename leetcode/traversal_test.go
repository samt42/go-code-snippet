//Generated Test_traversal
package leetcode

import "testing"

func Test_traversal(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			traversal(tt.args.node)
		})
	}
}
