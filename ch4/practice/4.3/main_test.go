package main

import "testing"

func TestReverseArr(t *testing.T) {
	tests := []struct {
		a    [32]byte
		want [32]byte
	}{
		{
			[32]byte{1, 2, 3},
			[32]byte{29: 3, 30: 2, 31: 1},
		},
	}

	for _, v := range tests {
		var b [32]byte
		b = v.a
		ReverseArr(&v.a)
		if v.a != v.want {
			t.Errorf("%v -> %v, want: %v", b, v.a, v.want)
		}
	}
}
