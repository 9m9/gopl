package slice

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

// 把一个 int slice 逆序
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 把一个 int slice rotate n 位
func Rotate(a []int, n int) {
	n = n % len(a)
	Reverse(a[:n])
	Reverse(a[n:])
	Reverse(a)
}
