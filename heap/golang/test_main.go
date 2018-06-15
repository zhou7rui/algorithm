package main

import (
	"fmt"
	"math/rand"

	"./heap"
)

func main() {
	integerMaxHeap := heap.MaxHeap{}
	integerMaxHeap.Init(15)
	for i := 0; i < 15; i++ {
		integerMaxHeap.Insert(heap.Integer(rand.Intn(20)))
	}
	fmt.Println(integerMaxHeap.GetData())
	fmt.Println(integerMaxHeap.ExtractMax())
	fmt.Println(integerMaxHeap.GetData())
}
