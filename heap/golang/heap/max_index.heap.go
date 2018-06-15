package heap

type MaxIndexHeap struct {
	data     []Comparable
	indexes  []int
	reverse  []int
	count    int
	capacity int
}

func (m *MaxIndexHeap) Init(capacity int) {

	m.data = make([]Comparable, capacity+1)
	m.indexes = make([]int, capacity+1)
	m.reverse = make([]int, capacity+1)
	for i := 0; i <= capacity; i++ {
		m.reverse[i] = 0
	}
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

	i += 1
	m.data[i] = item
	m.indexes[m.count+1] = i
	m.reverse[i] = m.count + 1
	m.count++
	m.shiftUp(m.count)

}

func (m *MaxIndexHeap) shiftUp(k int) {

	for k > 1 && m.data[m.indexes[k]].CompareTo(m.data[m.indexes[k/2]]) > 0 {
		m.indexes[k], m.indexes[k/2] = m.indexes[k/2], m.indexes[k]
		m.reverse[m.indexes[k/2]] = k / 2
		m.reverse[m.indexes[k]] = k

		k /= 2
	}
}

func (m *MaxIndexHeap) ExtractMax() Comparable {

	item := m.data[m.indexes[1]]
	m.indexes[1], m.indexes[m.count] = m.indexes[m.count], m.indexes[1]
	m.reverse[m.indexes[1]] = 1
	m.reverse[m.indexes[m.count]] = 0
	m.count--
	m.shiftDown(1)
	return item
}

func (m *MaxIndexHeap) ExtractMaxIndex() int {

	item := m.indexes[1] - 1
	m.indexes[1], m.indexes[m.count] = m.indexes[m.count], m.indexes[1]
	m.reverse[m.indexes[1]] = 1
	m.reverse[m.indexes[m.count]] = 0
	m.count--
	m.shiftDown(1)
	return item
}

func (m *MaxIndexHeap) GetItem(index int) Comparable {
	return m.data[index+1]
}

func (m *MaxIndexHeap) contain(i int) bool {
	if i > m.capacity {
		return false
	}

	return m.reverse[i+1] != 0
}

func (m *MaxIndexHeap) Change(i int, item Comparable) {
	if m.contain(i) {
		return
	}

	i += 1
	m.data[i] = item
	// 找到indexes[j] == i, i 表示data[i]在堆中的元素的位置
	// 之后shiftUp(j),再shiftDown(j)

	// for j := 1; j <= m.count; i++ {
	// 	if m.indexes[j] == i {
	// 		m.shiftUp(j)
	// 		m.shiftDown(j)
	// 	}
	// }

	j := m.reverse[i]
	m.shiftDown(j)
	m.shiftUp(j)

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
		m.reverse[m.indexes[k]] = k
		m.reverse[m.indexes[j]] = j
		k = j

	}

}
