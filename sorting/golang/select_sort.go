package main

import (
	"fmt"
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

func arrayPrint(arr []int) {
	fmt.Print("arr: [")
	isComma := false
	for i := 0; i < len(arr); i++ {
		if isComma {
			fmt.Print(", ")
		}
		fmt.Print(arr[i])
		isComma = true
	}
	fmt.Println("]")
}

func main() {
	arr := []int{1, 7, 8, 3, 5, 2, 6, 9, 4, 0}
	arr = selectSort(arr, len(arr))
	arrayPrint(arr)
}
