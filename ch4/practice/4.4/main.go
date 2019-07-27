package main

// 一次遍历就可以完成元素旋转
func Rotate1(a []int, n int) {
	if n > len(a) {
		n = n % len(a)
	}
	b := a[:n]
	copy(a[:len(a)-n], a[n:])
	copy(a[len(a)-n:], b)
}
