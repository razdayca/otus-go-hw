package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	len   int
}

func NewList() List {
	return new(list)
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		node := l.first
		l.first = node.Next
		l.len--
	} else if i.Next == nil {
		node1 := l.first
		var node2 *ListItem
		for node1.Next != nil {
			node2 = node1
			node1 = node1.Next
		}
		node2.Next = nil
		l.len--
	} else {
		node1 := l.first
		for i := 0; i < l.len-1; i++ {
			node1 = node1.Next
		}
		for i := l.len - 1; i > 0; i-- {
			node1 = node1.Prev
		}
		node2 := node1.Next
		node1.Next = node2.Next
		node1 = node2.Prev
		node2.Prev = node1.Prev
		l.len--
	}
}

func (l *list) PushBack(v interface{}) *ListItem {
	s := ListItem{
		Value: v,
		Prev:  l.last,
	}
	l.last.Next = &s
	l.last = &s
	l.len++
	return &s
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.first == nil {
		s := ListItem{
			Value: v,
		}
		l.first = &s
		l.last = &s
		l.len++
		return &s
	} else {
		s := ListItem{
			Value: v,
			Next:  l.first,
		}
		l.first.Prev = &s
		l.first = &s
		l.len++
		return &s
	}
}

func (l list) Back() *ListItem {
	return l.last
}

func (l list) Front() *ListItem {
	return l.first
}

func (l list) Len() int {
	return l.len
}
