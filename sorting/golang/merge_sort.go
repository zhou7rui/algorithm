package main

import (
	"./help"
)

func merge_sort(arr []int, l int, r int) {

	if l >= r {
		return
	}
	mid := l + (r-l)/2
	merge_sort(arr, l, mid)
	merge_sort(arr, mid+1, r)
	merge(arr, l, mid, r)
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

func sort(arr []int, n int) []int {
	merge_sort(arr, 0, n-1)
	return arr
}

func main() {
	arr := help.GenRanArray(1000000, 1, 999999)
	//help.ArrayPrint(arr)
	help.TestSort(sort)(arr, len(arr))
	//help.ArrayPrint(arr)

}
