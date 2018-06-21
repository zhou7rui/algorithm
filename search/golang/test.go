package main

import (
	"fmt"

	"./search"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(search.BinarySearch(arr, 6))
}
