package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "Case1", args: args{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{1, 0}},
		{name: "Case2", args: args{nums: []int{3, 2, 4}, target: 6}, want: []int{2, 1}},
		{name: "Case3", args: args{nums: []int{3, 3}, target: 6}, want: []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSumBruteForce(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "Case1", args: args{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{name: "Case2", args: args{nums: []int{3, 2, 4}, target: 6}, want: []int{1, 2}},
		{name: "Case3", args: args{nums: []int{3, 3}, target: 6}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSumBruteForce(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
