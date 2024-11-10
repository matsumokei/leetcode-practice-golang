package leetcode

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct{
		name string
		input int
		match bool
	}{
		{
			name: "zero",
			input: 0,
			match: true,
		},
		{
			name: "negative value",
			input: -121,
			match: false,
		},
		{
			name: "multiple of 10",
			input: 10,
			match: false,
		},
		{
			name: "symmetry1",
			input: 121,
			match: true,
		},
		{
			name: "symmetry2",
			input: 345121543,
			match: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isPalindrome1(tt.input) == true {
				fmt.Printf("input:%v  judge: true", tt.input)
			}else{
				fmt.Printf("input:%v  judge: false", tt.input)
			}
			fmt.Printf("\n\n")
			
		})
	}
}
