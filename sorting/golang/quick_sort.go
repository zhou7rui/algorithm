package main

import (
	"math/rand"

	"./help"
)

func partition(arr []int, l, r int) int {

	// 随机一个位置
	ran := rand.Intn((r - l)) + l
	arr[l], arr[ran] = arr[ran], arr[ran]

	temp := arr[l]

	j := l
	for i := l + 1; i < r+1; i++ {
		if temp > arr[i] {
			arr[j+1], arr[i] = arr[i], arr[j+1]
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func quickInsertSort(arr []int, l, r int) {

	for i := l + 1; i <= r; i++ {

		temp := arr[i]
		j := i
		for j > l && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp

	}

}

func quickSort(arr []int, l, r int) {

	if r-l <= 15 {
		quickInsertSort(arr, l, r)
		return
	}
	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

func QuickSort(arr []int, n int) []int {
	quickSort(arr, 0, n-1)
	return arr
}

func main() {

	arr := help.GenRanArray(1000000, 1, 99999)
	// help.ArrayPrint(arr)
	help.TestSort(QuickSort)(arr, len(arr))
	// help.ArrayPrint(arr)
}
