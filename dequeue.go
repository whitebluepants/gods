package main

type Dequeue[T any] struct {
	data []T
}

func (d *Dequeue[T]) PushBack(e T) {
	d.data = append(d.data, e)
}
func (d *Dequeue[T]) PushFront(e T) {
	d.data = append([]T{e}, d.data...)
}
func (d *Dequeue[T]) PopBack() T {
	e := d.data[len(d.data)-1]
	d.data = d.data[0 : len(d.data)-1]
	return e
}
func (d *Dequeue[T]) PopFront() T {
	e := d.data[0]
	d.data = d.data[1:]
	return e
}
func (d *Dequeue[T]) Front() T {
	return d.data[0]
}
func (d *Dequeue[T]) Back() T {
	return d.data[len(d.data)-1]
}
func (d *Dequeue[T]) Size() int {
	return len(d.data)
}
