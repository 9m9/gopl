package main

import (
	"abel-abel/gopl/ch4/equal"
	"testing"
)

func TestRotate1(t *testing.T) {
	tests := []struct {
		a    []int
		b    int
		want []int
	}{
		{
			[]int{1, 2, 3},
			1,
			[]int{2, 3, 1},
		},
		{
			[]int{1, 2, 3},
			2,
			[]int{3, 1, 2},
		},
		{
			[]int{1, 2, 3},
			3,
			[]int{1, 2, 3},
		},
	}
	for _, v := range tests {
		originA := make([]int, len(v.a), len(v.a))
		copy(originA, v.a)
		if Rotate1(v.a, v.b); !equal.Equal(v.a, v.want) {
			t.Errorf("%v rotate %d -> %v, want %v", originA, v.b, v.a, v.want)
		}
	}
}
