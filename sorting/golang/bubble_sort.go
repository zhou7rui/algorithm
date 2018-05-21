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


// 冒泡排序优化思路1
// 设置开关，当第二层循环不在进行交换时说明排序完成 提前结束外层循环
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

// 冒泡排序优化思路2
// 记录最后一次交换位置， 上层循环从最后一次位置开始
func bubbleSort2(arr []int, n int) []int {

	flag := n
	for flag > 0 {
		n = flag
		flag = 0
		for i := 1; i < n; i ++ {
			if arr[i] < arr[i - 1] {
				arr[i], arr[i - 1 ] = arr[i - 1], arr[i]
				flag = i
			}
		}

	}


	return nil

}

func main() {
	arr := help.GenRanArray(10000, 1, 9999)
	// help.ArrayPrint(arr)
	help.TestSort(bubbleSort2)(arr, len(arr))
}
