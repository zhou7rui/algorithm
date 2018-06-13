package main

import (
	"math/rand"

	"./help"
)

// 将 数组分成3部分  小于某个值的 等于某个值 大于某个值
//         <              =               <
// [| | | | | | | | |              | | | | |  | | |]
//  l               lt i          gt              r
//
//
func quick3Sort(arr []int, l, r int) {

	if r-l <= 15 {
		quick3InsertSort(arr, l, r)
		return
	}

	//partition
	// 随机一个位置
	ran := rand.Intn(r-l) + l
	arr[l], arr[ran] = arr[ran], arr[ran]
	temp := arr[l]

	lt := l     //arr[l + 1 ... lt] < v
	gt := r + 1 //arr[gt....r] > v
	i := l + 1  //arr[lt + 1 ... i) == v

	for i < gt {
		if arr[i] < temp {
			arr[lt+1], arr[i] = arr[i], arr[lt+1]
			lt++
			i++
		} else if arr[i] > temp {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else { // arr[i] == temp
			i++
		}
	}

	// 交换一下 位置后lt的位置便为等于temp的最小索引
	arr[l], arr[lt] = arr[lt], arr[l]

	quick3Sort(arr, l, lt-1)
	quick3Sort(arr, gt, r)

}

func quick3InsertSort(arr []int, l, r int) {

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

func Quick3Sort(arr []int, n int) []int {
	quick3Sort(arr, 0, n-1)
	return arr
}

func main() {

	arr := help.GenRanArray(10000, 1, 9999)
	// help.ArrayPrint(arr)
	help.TestSort(Quick3Sort)(arr, len(arr))
	// help.ArrayPrint(arr)
}
