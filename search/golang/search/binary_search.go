package search

import (
	"fmt"
)

// 递归实现
func BinaryRecSearch(arr []int, l, r, target int) int {

	mid := l + (r-l)/2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] > target {
		return BinaryRecSearch(arr, l, mid-1, target)
	} else {
		return BinaryRecSearch(arr, mid+1, r, target)
	}

}

//二分查找法 返回索引值
func BinarySearch(arr []int, target int) int {

	l := 0
	r := len(arr) - 1

	for l <= r {

		mid := l + (r-l)/2

		if arr[mid] == target {
			fmt.Println(0)
			return mid
		}
		if arr[mid] < target {
			l = mid + 1
		} else {

			r = mid - 1
		}
	}
	return -1

}
