package leetcode

import (
	"testing"
)

func TestReverseNodesInKGroups(t *testing.T) {
	var tests = []struct {
		before []int
		k      int
		after  []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{[]int{1}, 1, []int{1}},
	}

	for _, tt := range tests {
		after := reverseKGroup(NewList(tt.before), tt.k)
		if !EqualList(after, tt.after) {
			t.Errorf("reverseKGroup(%v, %v) return %v, want %v", tt.before, tt.k, after, tt.after)
		}
	}

}
