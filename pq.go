package main

type priorityQueue[T any] struct {
	data []T
}

func (pq *priorityQueue[T]) Len() int {
	return len(pq.data)
}
func (pq *priorityQueue[T]) Less(i, j int) bool {
	return cmp(pq.data[i], pq.data[j])
}
func cmp(x interface{}, y interface{}) bool {
	switch x.(type) {
	case pair:
		a, b := x.(pair), y.(pair)
		return a.nextTime < b.nextTime
	case pair2:
		a, b := x.(pair2), y.(pair2)
		return a.cnt > b.cnt
	}
	return true
}
func (pq *priorityQueue[T]) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}
func (pq *priorityQueue[T]) Push(x any) {
	pq.data = append(pq.data, x.(T))
}
func (pq *priorityQueue[T]) Pop() any {
	r := pq.data[pq.Len()-1]
	pq.data = pq.data[:pq.Len()-1]
	return r
}
func (pq *priorityQueue[T]) Top() any {
	r := pq.data[0]
	return r
}

type pair struct {
	cnt      int
	nextTime int
	name     byte
}
type pair2 struct {
	cnt      int
	nextTime int
	name     byte
}
