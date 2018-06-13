package main

import (
	"./help"
)

func shiftDown(arr []int, n, k int) {

	for 2*k+1 < n {

		j := 2 * k

		if j+1 < n && arr[j] < arr[j+1] {
			j++
		}
		if arr[k] >= arr[j] {
			break
		}

		arr[j], arr[k] = arr[k], arr[j]

		k = j

	}

}

func HeapSort(arr []int, n int) []int {

	for i := (n - 1) / 2; i >= 0; i-- {
		shiftDown(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		shiftDown(arr, i, 0)

	}
	return arr
}

func main() {

	arr := help.GenRanArray(1000, 1, 999)
	help.TestSort(HeapSort)(arr, len(arr))
}
