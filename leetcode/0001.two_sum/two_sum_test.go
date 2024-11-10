package leetcode

import (
	"fmt"
	"testing"

)



func TestProblem(t *testing.T) {

	tests := []struct{
		name string
		numbers []int
		target int
		output []int
	}{
		{
			name: "pattern1",
			numbers: []int{3, 2, 4},
			target: 6,
			output: []int{1,2},
		},
		{
			name: "pattern2",
			numbers: []int{3, 2, 4},
			target: 5,
			output: []int{0, 1},
		},
		{
			name: "pattern3",
			numbers: []int{0, 8, 7, 3, 3, 4, 2},
			target: 11,
			output: []int{1,3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("【input】:%v  output】:%v\n", tt, twoSum(tt.numbers, tt.target))
			fmt.Printf("\n\n\n")
		})
	}

}