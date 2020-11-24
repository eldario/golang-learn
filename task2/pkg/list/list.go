package list

import (
	"container/list"
)

type NewList struct {
	items list.List
}

func New() *NewList {
	return new(NewList)
}

func (l *NewList) Insert(value int) {
	for e := l.items.Front(); e != nil; e = e.Next() {
		if value < e.Value.(int) {
			l.items.InsertBefore(value, e)
			return
		}
	}

	l.items.PushBack(value)
}

func (l *NewList) Remove(value int) {
	for e := l.items.Front(); e != nil; e = e.Next() {
		if value == e.Value.(int) {
			l.items.Remove(e)
		}
	}
}

func (l *NewList) GetItems() []int {
	var result []int

	for e := l.items.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(int))

	}

	return result
}

func (l *NewList) GetMax() int {
	var back = l.items.Back()
	if back == nil {
		return 0
	}

	return l.items.Back().Value.(int)
}

func (l *NewList) GetMin() int {
	var front = l.items.Front()
	if front == nil {
		return 0
	}

	return l.items.Front().Value.(int)
}

func (l *NewList) Equals(listValue NewList) bool {
	if l.items.Len() != listValue.items.Len() {
		return false
	}

	for item1, item2 := l.items.Front(), listValue.items.Front(); item1 != nil; item1, item2 = item1.Next(), item2.Next() {
		if item1.Value != item2.Value {
			return false
		}
	}

	return true
}