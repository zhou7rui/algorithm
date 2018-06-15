package heap

type MaxIndexHeap struct {
	data     []Comparable
	indexes  []int
	count    int
	capacity int
}

func (m *MaxIndexHeap) Init(capacity int) {

	m.data = make([]Comparable, capacity+1)
	m.indexes = make([]int, capacity+1)
	m.count = 0
	m.capacity = capacity

}

func (m *MaxIndexHeap) Size() int {
	return m.count
}

func (m *MaxIndexHeap) IsEmpty() bool {
	return m.count == 0
}

func (m *MaxIndexHeap) Insert(i int, item Comparable) {

	m.data[i] = item
	m.indexes[m.count+1] = i
	m.count++
	m.shiftUp(m.count)

}

func (m *MaxIndexHeap) shiftUp(k int) {

	for k > 1 && m.data[m.indexes[k]].CompareTo(m.data[m.indexes[k/2]]) > 0 {
		m.indexes[k], m.indexes[k/2] = m.indexes[k/2], m.indexes[k]
		k /= 2
	}
}

func (m *MaxIndexHeap) ExtractMax() Comparable {

	item := m.data[m.indexes[1]]
	m.indexes[1], m.indexes[m.count] = m.indexes[m.count], m.indexes[1]
	m.count--
	m.shiftDown(1)
	return item
}

func (m *MaxIndexHeap) ExtractMaxIndex() int {

	item := m.indexes[1] - 1
	m.indexes[1], m.indexes[m.count] = m.indexes[m.count], m.indexes[1]
	m.count--
	m.shiftDown(1)
	return item
}

func (m *MaxIndexHeap) GetItem(index int) Comparable {
	return m.data[index+1]
}

func (m *MaxIndexHeap) Change(i int, item Comparable) {
	i += 1
	m.data[i] = item
	// 找到indexes[j] == i, i 表示data[i]在堆中的元素的位置
	// 之后shiftUp(j),再shiftDown(j)

	for j := 1; j <= m.count; i++ {
		if m.indexes[j] == i {
			m.shiftUp(j)
			m.shiftDown(j)
		}

	}

}

func (m *MaxIndexHeap) shiftDown(k int) {
	for 2*k <= m.count {
		j := k * 2
		if j+1 <= m.count && m.data[m.indexes[j]].CompareTo(m.data[m.indexes[j+1]]) < 0 {
			j++
		}
		if m.data[m.indexes[k]].CompareTo(m.data[m.indexes[j]]) >= 0 {
			break
		}
		m.indexes[k], m.indexes[j] = m.indexes[j], m.indexes[k]
		k = j

	}

}
