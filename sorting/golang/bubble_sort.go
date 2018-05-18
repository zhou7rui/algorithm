package main

import "./help"

func bubbleSort(arr []int, n int) []int {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func main() {
	arr := []int{1, 7, 8, 3, 5, 2, 6, 9, 4, 0}
	help.ArrayPrint(bubbleSort(arr, len(arr)))
}
