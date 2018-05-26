package main

import (
	"./help"
)

func InsertSort(arr []int, n int) []int {
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i
		for j > 0 && arr[j-1] > temp {
			arr[j] = arr[j-1]
			j = j - 1
		}
		arr[j] = temp
	}
	return arr
}

func InsertSwapSort(arr []int, n int) []int {
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}

	}
	return arr
}

func main() {
	arr := help.GenRanArray(10000, 1, 9999)
	// help.ArrayPrint(arr)
	help.TestSort(InsertSort)(arr, len(arr))

}
