package help

import "fmt"

func ArrayPrint(arr []int) {
	fmt.Print("arr: [")
	isComma := false
	for i := 0; i < len(arr); i++ {
		if isComma {
			fmt.Print(", ")
		}
		fmt.Print(arr[i])
		isComma = true
	}
	fmt.Println("]")
}
