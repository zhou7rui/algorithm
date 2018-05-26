package main

import (
	"math/rand"

	"./help"
)

func quick2Sort(arr []int, l, r int) {

	if r-l <= 15 {
		quick2InsertSort(arr, l, r)
		return
	}
	p := partition2(arr, l, r)
	quick2Sort(arr, l, p-1)
	quick2Sort(arr, p+1, r)

}

func partition2(arr []int, l, r int) int {
	// 随机一个位置
	ran := rand.Intn(r-l+1) + l
	arr[l], arr[ran] = arr[ran], arr[ran]
	temp := arr[l]
	i, j := l+1, r
	for true {
		for i <= r && arr[i] < temp {
			i++
		}
		for j >= l+1 && arr[j] > temp {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[j], arr[l] = arr[l], arr[j]
	return j
}

func quick2InsertSort(arr []int, l, r int) {

	for i := l + 1; i < r+1; i++ {

		temp := arr[i]
		j := i
		for j > l && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp

	}

}

func Quick2Sort(arr []int, n int) []int {
	quick2Sort(arr, 0, n-1)
	return arr
}

func main() {
	arr := help.GenRanArray(1000000, 1, 999999)
	help.TestSort(Quick2Sort)(arr, len(arr))
}
