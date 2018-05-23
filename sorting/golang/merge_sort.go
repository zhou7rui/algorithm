package main

import "./help"


func merge_sort(arr []int,l int, r int) []int{
	
	if l >= r{
		return nil
	}
	mid := r + (l - r) / 2
	merge_sort(arr, l, mid + 1)
	merge_sort(arr,mid, r)
	arr = merge(arr, l, mid, r)
	return arr
}


func merge(arr []int, l int, mid int, r int)[] int{

	aux := arr
	i, j := l, mid + 1
	for k := i; k <= r; k ++ {
		switch {
			case i > mid: 
				arr[k] =  aux[j - l]
				j ++
			case j > r:
				arr[k] = aux[i - l]
				i ++
			case aux[i - l] > aux[j - l] :
				arr[k] = arr[i - l]
				i ++
			default:
				arr[k] = arr[j - l]
				j ++ 

		}
	}
	return arr

}

func sort(arr []int, n int) []int{
	merge_sort(arr, 0, n - 1)
	return nil
}
func main(){

	arr := help.GenRanArray(10000, 1, 9999)
	help.TestSort(sort)(arr, len(arr))
}


