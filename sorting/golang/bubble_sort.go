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


func bubbleSort1(arr []int, n int) []int {
	for i := 0; i < n-1; i++ {
		flag := true
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return arr
}


func main() {
	arr := help.GenRanArray(10000, 1, 9999)
	// help.ArrayPrint(arr)
	help.TestSort(bubbleSort1)(arr, len(arr))
}
