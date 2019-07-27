package main

import "fmt"

func main() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2, summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", q)
			}
		}
	}

	fmt.Println(months[12])
	fmt.Printf("summer=%v, len=%d, cap=%d, summer[2]=%s\n",
		summer, len(summer), cap(summer), summer[2])
	// out of range
	fmt.Printf("summer[3:]=%v\n", summer[3:])
	fmt.Printf("summer[3:5]=%v\n", summer[3:5])
	fmt.Printf("summer=%v\n", summer)
	// fmt.Printf("summer[3]=%s\n", summer[3])

	printSliceLenCap(summer)
	printSliceLenCap(summer[3:])
	printSliceLenCap([]string{})
	var s []string
	printSliceLenCap(s)
	s = nil
	printSliceLenCap(s)
	s = []string(nil)
	printSliceLenCap(s)
	s = []string{}
	printSliceLenCap(s)

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	xcap := cap(x)

	if xcap >= zlen {
		z = x[:zlen]
		z[zlen-1] = y
		return z
	}

	z = make([]int, zlen, xcap*2)
	copy(z, x)
	z[zlen-1] = y
	return z
}

func printSliceLenCap(s []string) {
	fmt.Printf("s=%v, len=%d, cap=%d, isNil=%t\n", s, len(s), cap(s), s == nil)
}

func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Rotate(a []int, n int) {
	n = n % len(a)
	Reverse(a[:n])
	Reverse(a[n:])
	Reverse(a)
}
