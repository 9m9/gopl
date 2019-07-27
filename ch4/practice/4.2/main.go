package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var num = flag.Int("num", 256, "sha method")

func main() {
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		strByte := []byte(str)
		if *num == 256 {
			sh := sha256.Sum256(strByte)
			fmt.Printf("%q\t-> %x\n", str, sh)
		} else if *num == 384 {
			sh := sha512.Sum384(strByte)
			fmt.Printf("%q\t-> %x\n", str, sh)
		} else if *num == 512 {
			sh := sha512.Sum512(strByte)
			fmt.Printf("%q\t-> %x\n", str, sh)
		}
	}
}
