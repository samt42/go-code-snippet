//Generated TestCheckPossibility
//Generated Test_countPeak
package leetcode

import (
	"testing"
)

func TestCheckPossibility(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name:"1", args:args{nums:[]int{3,2,1}}, want:false},
		{name:"2", args:args{nums:[]int{3,2,1,1}}, want:false},
		{name:"3", args:args{nums:[]int{3,2,2,1}}, want:false},
		{name:"4", args:args{nums:[]int{3,2,2,1,1}}, want:false},
		{name:"5", args:args{nums:[]int{2,2,2,1,1}}, want:false},
		{name:"6", args:args{nums:[]int{2,1,1,1,1}}, want:true},
		{name:"7", args:args{nums:[]int{1,1,1}}, want:true},
		{name:"8", args:args{nums:[]int{1,1,1,1}}, want:true},
		{name:"9", args:args{nums:[]int{1,1,2,1}}, want:true},
		{name:"10", args:args{nums:[]int{3,3,1}}, want:true},
		{name:"11", args:args{nums:[]int{1,2,3}}, want:true},
		{name:"12", args:args{nums:[]int{1,2,1}}, want:true},
		{name:"13", args:args{nums:[]int{-1,4,2,3}}, want:true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.args.nums); got != tt.want {
				t.Errorf("CheckPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
