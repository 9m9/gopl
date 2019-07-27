package main

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
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
		if Rotate(v.a, v.b); !Equal(v.a, v.want) {
			t.Errorf("%v rotate %d -> %v, want %v", originA, v.b, v.a, v.want)
		}
	}
}

func TestAppendInt(t *testing.T) {
	tests := []struct {
		a    []int
		b    int
		want []int
	}{
		{
			[]int{1, 2, 3},
			4,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 2},
			3,
			[]int{1, 2, 3},
		},
		{
			[]int{1},
			2,
			[]int{1, 2},
		},
	}

	for _, v := range tests {
		if got := appendInt(v.a, v.b); !Equal(got, v.want) {
			fmt.Printf("%v append %d get %v, want %v", v.a, v.b, got, v.want)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		a    []int
		want []int
	}{
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{100, 200},
			[]int{200, 100},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{4, 3, 2, 1},
		},
	}
	for _, v := range tests {
		var originA []int = make([]int, len(v.a), len(v.a))
		copy(originA, v.a)
		// fmt.Printf("originA=%v, v.a=%v, v.want=%v\n", originA, v.a, v.want)
		if Reverse(v.a[:]); !Equal(v.a, v.want) {
			t.Errorf("%v reversed: %v, want %q", originA, v.a, v.want)
		}
		// fmt.Printf("originA=%v, v.a=%v, v.want=%v\n", originA, v.a, v.want)
	}
}
