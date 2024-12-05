package main

import (
	"fmt"
)

// ReverseSliceInPlace 泛型函数，用于原地翻转任何类型的 slice
func ReverseSliceInPlace[T any](s []T) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}

func main() {
	// 测试整数 slice
	intSlice := []int{1, 2, 3, 4, 5}
	ReverseSliceInPlace(intSlice)
	fmt.Println("Reversed int slice (in place):", intSlice)

	// 测试字符串 slice
	strSlice := []string{"a", "b", "c", "d"}
	ReverseSliceInPlace(strSlice)
	fmt.Println("Reversed string slice (in place):", strSlice)

	// 测试浮点数 slice
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	ReverseSliceInPlace(floatSlice)
	fmt.Println("Reversed float slice (in place):", floatSlice)
}
