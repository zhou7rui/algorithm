package main

import (
	"./help"
)

func selectSort(arr []int, n int) []int {
	for i := 0; i < n; i++ {
		minIndex := i
		// 遍历出最小的索引位置
		for j := i + 1; j < n; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		// swap
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}

func main() {
	arr := help.GenRanArray(10000, 1, 9999)
	// help.ArrayPrint(arr)
	help.TestSort(selectSort)(arr, len(arr))
}
