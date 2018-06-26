package main

import (
	"fmt"

	"./search"
)

func main() {

	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// fmt.Println(search.BinarySearch(arr, 6))

	bst := search.BST{}

	bst.Insert(11, "AA")

	bst.Insert(55, "BB")

	bst.Insert(88, "CC")

	bst.Insert(77, "DD")

	bst.Insert(99, "EE")

	bst.Insert(66, "FF")

	bst.Insert(64, "GG")

	fmt.Println(bst.Max())
}
