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

// 移除第 i 个元素，并保留剩余元素的顺序
func Remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

// 移除第 i 个元素，不需要保留剩余元素的顺序
func Remove1(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func NonEmpty(s []string) []string {
	j := 0
	for _, v := range s {
		if v != "" {
			s[j] = v
			j++
		}
	}
	return s[:j]
}

// 重写针对 int slice 的 append 函数
func AppendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	xcap := cap(x)

	if xcap >= zlen {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < xcap*2 {
			zcap = xcap * 2
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	copy(z[len(x):], y)
	return z
}

func printSliceLenCap(s []string) {
	fmt.Printf("s=%v, len=%d, cap=%d, isNil=%t\n", s, len(s), cap(s), s == nil)
}

// 把一个 int slice 逆序
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 判断两个 int slice 是否相等
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

func StringSliceEqual(a []string, b []string) bool {
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

// 把一个 int slice rotate n 位
func Rotate(a []int, n int) {
	n = n % len(a)
	Reverse(a[:n])
	Reverse(a[n:])
	Reverse(a)
}
