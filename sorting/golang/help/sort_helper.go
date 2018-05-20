package help

import (
	"fmt"
	"runtime"
	"reflect"
	"math/rand"
	"time"
)

func TestSort(f func([]int,int) []int) func([]int,int) {
	
	return func(arr []int, n int){
		startTime := time.Now()
		p := f(arr, n)
		endTime := time.Now()
		if !isSort(p) {
			fmt.Println("sort failure")
			return
		}
		fmt.Printf("%s runTime %d ms \n",getFunctionName(f), endTime.Sub(startTime) / 1e6)
	}

}

func GenRanArray(count, start, end int) []int{
	  //范围检查  
	  if end <= start {  
        return nil  
    }  
    //存放结果的slice  
    nums := make([]int, 0)  
    //随机数生成器，加入时间戳保证每次生成的随机数不一样  
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for len(nums) < count {  
        //生成随机数  
        num := r.Intn((end - start)) + start  
		
        //查重  
        // exist := false  
        // for _, v := range nums {  
        //     if v == num {  
        //         exist = true  
        //         break  
        //     }  
        // }  
        nums = append(nums, num)  
        
    }
	return nums
}


func isSort(arr []int) bool{
	for i := 0; i < len(arr) - 1; i++{
		if arr[i] > arr[i + 1]{
			return false
		}
	}
	return true
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
  }

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
