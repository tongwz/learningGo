package main

import (
	"fmt"
)

// 冒泡排序
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 归并排序
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

// 快速排序
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	left, right := []int{}, []int{}

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

// 选择排序
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// 插入排序
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("原始数组:", arr)

	bubbleSort(arr)
	fmt.Println("冒泡排序结果:", arr)

	arr = []int{90, 64, 34, 25, 12, 22, 11}

	sortedArr := mergeSort(arr)
	fmt.Println("归并排序结果:", sortedArr)

	arr = []int{64, 34, 25, 12, 22, 11, 90}

	sortedArr = quickSort(arr)
	fmt.Println("快速排序结果:", sortedArr)

	arr = []int{64, 34, 25, 12, 22, 11, 90}

	selectionSort(arr)
	fmt.Println("选择排序结果:", arr)

	arr = []int{64, 34, 25, 12, 22, 11, 90}

	insertionSort(arr)
	fmt.Println("插入排序结果:", arr)
}
