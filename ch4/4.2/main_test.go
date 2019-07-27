package main

import (
	"abel-abel/gopl/ch4/equal"
	"abel-abel/gopl/ch4/slice"
	"fmt"
	"testing"
)

func TestNonEmpty(t *testing.T) {
	tests := []struct {
		a    []string
		want []string
	}{
		{
			[]string{"a", "", "b"},
			[]string{"a", "b"},
		},
	}

	for _, v := range tests {
		if got := slice.NonEmpty(v.a); !equal.StringSliceEqual(got, v.want) {
			t.Errorf("%q -> %q, want: %q", v.a, got, v.want)
		}
	}
}

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
		if slice.Rotate(v.a, v.b); !equal.Equal(v.a, v.want) {
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
		if got := slice.AppendInt(v.a, v.b); !equal.Equal(got, v.want) {
			fmt.Printf("%v append %d get %v, want %v", v.a, v.b, got, v.want)
		}
	}

	var x, y []int
	for i := 0; i < 10; i++ {
		y = slice.AppendInt(x, i)
		fmt.Printf("i=%d cap=%d y=%d\n", i, cap(y), y)
		x = y
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
		if slice.Reverse(v.a[:]); !equal.Equal(v.a, v.want) {
			t.Errorf("%v reversed: %v, want %q", originA, v.a, v.want)
		}
		// fmt.Printf("originA=%v, v.a=%v, v.want=%v\n", originA, v.a, v.want)
	}
}
