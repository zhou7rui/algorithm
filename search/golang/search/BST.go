package search

import (
	"fmt"
	"reflect"
)

type Comparable interface {
	CompareTo(t Comparable) int
}

type Key int

func (k Key) CompareTo(t Key) int {
	var anotherKey Key
	// 通过使用反射获取值并进行转型
	anotherKey = reflect.ValueOf(t).Interface().(Key)
	return compare(k, anotherKey)
}

// 比较
func compare(x, y Key) int {
	if x < y {
		return -1
	} else if x > y {
		return 1
	} else {
		return 0
	}
}

type Node struct {
	key   Key
	value string
	left  *Node
	right *Node
}

func (n *Node) init(key Key, value string) {
	n.key = key
	n.value = value
}

type BST struct {
	root  *Node
	count int
}

func (b *BST) Insert(key Key, value string) {
	b.root = b.insert(b.root, key, value)
}

func (b *BST) insert(node *Node, key Key, value string) *Node {

	if node == nil {
		b.count++
		newNode := Node{}
		newNode.init(key, value)
		return &newNode
	}

	if node.key.CompareTo(key) == 0 {
		node.key = key
	} else if node.key.CompareTo(key) < 0 {
		node.right = b.insert(node.right, key, value)
	} else { //node.key.CompareTo(key) > 0
		node.left = b.insert(node.left, key, value)
	}

	return node

}

func (b *BST) Search(key Key) *string {

	return b.search(b.root, key)
}

func (b *BST) search(node *Node, key Key) *string {

	if node == nil {
		return nil
	}
	if node.key == key {
		return &(node.value)
	} else if node.key.CompareTo(key) > 0 {
		return b.search(node.left, key)
	} else { //node.key.CompareTo(key) < 0
		return b.search(node.right, key)
	}

}

func (b *BST) Contain(key Key) {
	b.contain(b.root, key)
}

func (b *BST) contain(node *Node, key Key) bool {

	if node == nil {
		return false
	}

	if node.key == key {
		return true
	} else if node.key.CompareTo(key) < 0 {
		return b.contain(node.left, key)
	} else {
		return b.contain(node.right, key)
	}
}

// 前序遍历
func (b *BST) PreOrder() {
	preOrder(b.root)
}

// 中序遍历
func (b *BST) InOrder() {
	inOrder(b.root)
}

//后序遍历
func (b *BST) PostOrder() {
	postOrder(b.root)
}

func preOrder(node *Node) {

	if node != nil {
		fmt.Println(node.key)
		preOrder(node.left)
		preOrder(node.right)
	}
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.left)
		fmt.Println(node.key)
		inOrder(node.right)
	}
}

func postOrder(node *Node) {
	if node != nil {
		postOrder(node.left)
		postOrder(node.right)
		fmt.Println(node.key)
	}
}
