package remove

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
