package search

import (
	"fmt"
)

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
