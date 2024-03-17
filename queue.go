package main

type queue[T any] struct {
	data []T
}

func (d *queue[T]) push(e T) {
	d.data = append(d.data, e)
}
func (d *queue[T]) pop() T {
	e := d.data[0]
	d.data = d.data[1:]
	return e
}
func (d *queue[T]) top() T {
	return d.data[0]
}
func (d *queue[T]) size() int {
	return len(d.data)
}
