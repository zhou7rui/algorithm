package heap

type MaxHeap struct {
	Data  []interface{}
	Count int
}

func (maxHeap *MaxHeap) Init(capacity int) {
	maxHeap.Data = make([]interface{}, capacity+1)
	maxHeap.Count = 0
}

func (maxHeap *MaxHeap) Size() int {
	return maxHeap.Count
}

func (maxHeap *MaxHeap) IsEmpty() bool {
	return maxHeap.Count == 0
}

func main() {

}
