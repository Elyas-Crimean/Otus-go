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
	// Remove me after realization.
	// Place your code here.
	size  int
	front *ListItem
	back  *ListItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if l.front == i {
		l.front = i.Next
	}
	if l.back == i {
		l.back = i.Prev
	}
	l.size--
}

func (l *list) PushBack(v interface{}) *ListItem {
	i := &ListItem{
		Value: v,
		Prev:  l.back,
	}
	if l.back != nil {
		l.back.Next = i
	}
	if l.front == nil { // для пустого
		l.front = i
	}
	l.back = i
	l.size++
	return i
}

func (l *list) PushFront(v interface{}) *ListItem {
	i := &ListItem{
		Value: v,
		Next:  l.front,
	}
	if l.front != nil {
		l.front.Prev = i
	}
	if l.back == nil { // для пустого
		l.back = i
	}
	l.front = i
	l.size++
	return i
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}
	if i == l.back {
		l.back = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	i.Prev.Next = i.Next

	i.Next = l.front
	l.front.Prev = i
	l.front = i
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Len() int {
	return l.size
}

func NewList() List {
	return new(list)
}
