package main

import (
	"fmt"

	"./help"
)

// func shiftDown(arr []int, k int) {
// 	for k < len(arr) {
// 		j := k * 2
// 		if arr[j+1] > arr[j] {
// 			j++
// 		}
// 		if arr[k] >= arr[j] {
// 			break
// 		}
// 		arr[j], arr[k] = arr[k], arr[j]
// 		k = j
// 	}

// }

// func HeapSort(arr []int, n int) {

// }

func main() {
	fmt.Println("hello world ")
	arr := help.GenRanArray(10, 1, 9)
	maxIntheap := MaxIntHeap{}
	maxIntheap.Init(arr)

	for i := len(arr); i > 0; i-- {
		arr[i] = maxIntheap.ExtractMax()
	}
	help.ArrayPrint(arr)
}
