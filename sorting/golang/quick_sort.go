package main

import (
	"math/rand"

	"./help"
)

// 将数组分割成小于某个值的一部分 和 大于某个值的一部分
func partition(arr []int, l, r int) int {

	// 随机一个位置
	ran := rand.Intn((r - l)) + l
	arr[l], arr[ran] = arr[ran], arr[ran]
	temp := arr[l]

	j := l // arr[l + 1... j-1] < temp
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

/*
* 快速排序
* 将数组分割成小于某个值的一部分 和 大于某个值的一部分
* 分成的部分继续分割 达到排序的目的
 */
func quickSort(arr []int, l, r int) {

	// 当数据量小于15时采用插入排序的方式进行排序
	if r-l <= 15 {
		quickInsertSort(arr, l, r)
		return
	}

	// 寻找分割的位置
	p := partition(arr, l, r)
	// 继续分割
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
