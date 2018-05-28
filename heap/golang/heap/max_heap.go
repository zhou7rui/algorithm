package main

/*
*
*  最大堆
*
 */
import (
	"fmt"
	"math/rand"
	"reflect"
)

// Comparable 接口
type Comparable interface {
	CompareTo(t Comparable) int
}

// 定义结构体
type MaxHeap struct {
	data     []Comparable
	count    int
	capacity int
}

func (maxHeap *MaxHeap) Init(capacity int) {
	maxHeap.data = make([]Comparable, capacity+1)
	maxHeap.count = 0
	maxHeap.capacity = capacity
}

func (maxHeap *MaxHeap) Size() int {
	return maxHeap.count
}

func (maxHeap *MaxHeap) IsEmpty() bool {
	return maxHeap.count == 0
}

// 新增
func (maxHeap *MaxHeap) Insert(item Comparable) {
	if maxHeap.count+1 > maxHeap.capacity {
		fmt.Println("out of bounds")
		return
	}
	maxHeap.data[maxHeap.count+1] = item
	maxHeap.count++
	maxHeap.shiftUp(maxHeap.count)
}

func (maxHeap *MaxHeap) shiftUp(k int) {

	for k > 1 && maxHeap.data[k/2].CompareTo(maxHeap.data[k]) < 0 {
		maxHeap.data[k/2], maxHeap.data[k] = maxHeap.data[k], maxHeap.data[k/2]
		k /= 2
	}
}

// 提取出值
func (maxHeap *MaxHeap) ExtractMax() Comparable {
	item := maxHeap.data[1]
	maxHeap.data[1], maxHeap.data[maxHeap.count] = maxHeap.data[maxHeap.count], maxHeap.data[1]
	maxHeap.count--
	maxHeap.shitDown(1)
	return item
}
func (maxHeap *MaxHeap) shitDown(k int) {
	fmt.Println(k)
	for 2*k <= maxHeap.count {
		j := k * 2

		if j+1 < maxHeap.count && maxHeap.data[j+1].CompareTo(maxHeap.data[j]) > 0 {
			j++
		}
		if maxHeap.data[k].CompareTo(maxHeap.data[j]) >= 0 {
			break
		}
		maxHeap.data[j], maxHeap.data[k] = maxHeap.data[k], maxHeap.data[j]
		k = j

	}
}

// 包装类
type Integer int

// 实现Comparable  接口判断大小
func (integer Integer) CompareTo(t Comparable) int {
	var anotherInteger Integer
	// 通过使用反射获取值并进行转型
	anotherInteger = reflect.ValueOf(t).Interface().(Integer)
	return compare(integer, anotherInteger)
}

// 比较
func compare(x, y Integer) int {
	if x < y {
		return -1
	} else if x > y {
		return 1
	} else {
		return 0
	}

}

func main() {
	integerMaxHeap := MaxHeap{}
	integerMaxHeap.Init(15)
	for i := 0; i < 15; i++ {
		integerMaxHeap.Insert(Integer(rand.Intn(20)))
	}
	fmt.Println(integerMaxHeap.data)
	fmt.Println(integerMaxHeap.ExtractMax())
	fmt.Println(integerMaxHeap.data)
}
