package main

import (
	"./help"
)

func merge_sort(arr []int, l int, r int) {

	if  r - l <= 15 {
		insertSort(arr, l, r)
		return
	}
	mid := l + (r-l)/2
	merge_sort(arr, l, mid)
	merge_sort(arr, mid+1, r)
	if(arr[mid] > arr [mid+1]){
		merge(arr, l, mid, r)
	}
	
}

func merge(arr []int, l int, mid int, r int) {
	// copy 临时array
	aux := make([]int, r-l + 1)
	copy(aux,arr[l : r+1])

	i, j := l, mid+1
	for k := l; k < r+1; k++ {
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}

}

func insertSort(arr []int, l int, r int){

	for i:= l + 1; i < r + 1; i ++ {

		temp := arr[i]
		var j int
		for j = i; j > l && arr[j - 1] > temp; j -- {
			arr[j] = arr [j - 1]
		}
		arr[j] = temp
	}
}

func sort(arr []int, n int) []int {
	merge_sort(arr, 0, n-1)
	return arr
}

func min(x, y int) int{
	if x > y {
		return y
	}
	return x
}

// 至底向上归并排序过程
//  1 2   3 4    5 5  6 7   7 8   9  ^  
//  5 5   6 8    2 3  7 7   1 4 / 9  | 
//  5 6   5 8 /  2 7  3 7 / 1 4   9  | 
//  5 6 / 8 5 / 7 2 / 3 7 / 1 4 / 9  |
func mergeSortBU(arr []int, n int)[] int{

	for sz := 1; sz <= n; sz += sz {
		for i := 0; i + sz < n; i += sz * 2{
			merge(arr,i, i + sz - 1 , min(i + sz + sz -1,n - 1) )
		}

	}

	return arr


}

func main() {
	arr := help.GenRanArray(100000, 1, 99999)
	// help.ArrayPrint(arr)
	help.TestSort(mergeSortBU)(arr, len(arr))
	// help.ArrayPrint(arr)

}
