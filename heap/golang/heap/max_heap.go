package heap

/*
*
*  最大堆
*
 */
import (
	"fmt"
)

// 定义结构体
type MaxHeap struct {
	data     []Comparable
	count    int
	capacity int
}

func (maxHeap *MaxHeap) GetData() []Comparable {
	return maxHeap.data
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
	maxHeap.data = maxHeap.data[0:maxHeap.count]
	maxHeap.count--
	maxHeap.shiftDown(1)
	return item
}
func (maxHeap *MaxHeap) shiftDown(k int) {
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
