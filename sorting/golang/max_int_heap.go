package main

import (
	"fmt"

	"./help"
)

type MaxIntHeap struct {
	data     []int
	count    int
	capacity int
}

func (heap *MaxIntHeap) InsertSort(d int) {
	if heap.capacity < heap.count+1 {
		fmt.Println("out of bounds")
		return
	}

	heap.data[heap.count+1] = d
	heap.count++
	heap.shfitUp(heap.count)

}

func (heap *MaxIntHeap) shfitUp(k int) {
	for k > 1 && heap.data[k] > heap.data[k/2] {
		heap.data[k], heap.data[k/2] = heap.data[k/2], heap.data[k]
		k /= 2
	}

}

func (heap *MaxIntHeap) Size() int {
	return heap.count

}

func (heap *MaxIntHeap) Init(arr []int) {

	n := len(arr)
	heap.capacity = n
	heap.count = n

	heap.data = make([]int, n+1)
	for i := 0; i < n; i++ {
		heap.data[i+1] = arr[i]
	}

	for i := len(arr) / 2; i > 0; i-- {
		heap.shiftDown(i)
	}

}

func (heap *MaxIntHeap) shiftDown(k int) {
	for 2*k <= heap.count {
		j := k * 2
		if j+1 <= heap.count && heap.data[j+1] > heap.data[j] {
			j++
		}
		if heap.data[k] > heap.data[j] {
			break
		}
		heap.data[j], heap.data[k] = heap.data[k], heap.data[j]
		k = j
	}

}

func (heap *MaxIntHeap) ExtractMax() int {
	item := heap.data[1]
	heap.data[1], heap.data[heap.count] = heap.data[heap.count], heap.data[1]
	heap.data = heap.data[0:heap.count]
	heap.count--
	heap.shiftDown(1)
	return item
}

func HeapSort(arr []int, n int) []int {
	maxIntHeap := MaxIntHeap{}
	maxIntHeap.Init(arr)
	for i := n - 1; i >= 0; i-- {
		arr[i] = maxIntHeap.ExtractMax()

	}
	return arr

}

func main() {

	arr := help.GenRanArray(15, 1, 9)
	fmt.Println(arr)
	help.TestSort(HeapSort)(arr, len(arr))
	fmt.Println(arr)
	// arr := make([]int, 6)
	// fmt.Println(arr)

}
