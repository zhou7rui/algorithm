package heap

import "reflect"

// Comparable 接口
type Comparable interface {
	CompareTo(t Comparable) int
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
