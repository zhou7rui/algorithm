package main

import (
	"fmt"

	"./help"
)

func insertSort(arr []int, n int) []int {
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

func insertSwapSort(arr []int, n int) []int {
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
	arr := []int{1, 7, 8, 3, 5, 2, 6, 9, 4, 0}
	insertArr := insertSort(arr, len(arr))
	fmt.Print("insertSort")
	help.ArrayPrint(insertArr)
	insertSwapArr := insertSwapSort(arr, len(arr))
	fmt.Print("insertSwapSort")
	help.ArrayPrint(insertSwapArr)

}
