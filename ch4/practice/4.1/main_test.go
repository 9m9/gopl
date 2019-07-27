package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestDiffByteCountSum256(t *testing.T) {
	tests := []struct {
		a [32]byte
		b [32]byte
	}{
		{
			sha256.Sum256([]byte("x")),
			sha256.Sum256([]byte("y")),
		},
	}

	for _, v := range tests {
		got := diffByteCountSum256(v.a, v.b)
		fmt.Printf("%x %x -> %d\n", v.a, v.b, got)
	}
}

func TestDiffByteCount(t *testing.T) {
	tests := []struct {
		a    byte
		b    byte
		want byte
	}{
		{2, 3, 1},
		{1, 2, 2},
		{3, 3, 0},
		{5, 1, 1},
		{8, 7, 4},
	}

	for _, v := range tests {
		if got := diffByteCount(v.a, v.b); got != v.want {
			t.Errorf("diffByteCount(%d, %d) returns %d, want %d", v.a, v.b, got, v.want)
		}
	}
}
