package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")

	x := []byte("x")
	y := []byte("X")
	x1 := sha256.Sum256(x)
	y1 := sha256.Sum256(y)
	fmt.Printf("%x\n%x\n%T\n", x1, y1, x1)
}

func diffByteCountSum256(x [32]byte, y [32]byte) (cnt int) {
	for i := 0; i < len(x); i++ {
		cnt += int(diffByteCount(x[i], y[i]))
	}
	return
}

func diffByteCount(x byte, y byte) (cnt byte) {
	var k byte
	for i := 0; i < 8; i++ {
		k = 1 << uint8(i)
		if x&k != y&k {
			cnt++
		}
	}
	return
}

func byteCount(x byte) (cnt byte) {
	for i := 0; i < 8; i++ {
		if x&(1<<uint8(i)) != 0 {
			cnt++
		}
	}
	return cnt
}
