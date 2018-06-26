package search

import (
	"fmt"
)

type Queue interface {
	Push(n *Node)  // 添加
	Poll() *Node   // 移除最前面的元素
	Clear() bool   // 清除队列
	Size() int     // 对列的大小
	IsEmpty() bool // 队列是否为空
}

type sliceEntry struct {
	nodes []*Node
}

func NewQueue() *sliceEntry {
	return &sliceEntry{}
}

func (entry *sliceEntry) Push(node *Node) {
	entry.nodes = append(entry.nodes, node)
}

func (entry *sliceEntry) Poll() *Node {

	if entry.IsEmpty() {
		fmt.Println("queue is empty！")
		return &Node{}
	}

	firstNode := entry.nodes[0]
	entry.nodes = entry.nodes[1:]

	return firstNode
}

func (entry *sliceEntry) IsEmpty() bool {
	return len(entry.nodes) == 0
}

func (entry *sliceEntry) Size() int {
	return len(entry.nodes)
}

func (entry *sliceEntry) Clear() bool {

	if entry.IsEmpty() {
		fmt.Println("queue is empty！")
		return false
	}

	for i := 0; i < entry.Size(); i++ {
		entry.nodes[i] = &Node{}
	}
	entry.nodes = nil
	return true
}
