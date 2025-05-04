package main

import "fmt"

const MAX_SIZE = 5

type Queue struct {
	a [MAX_SIZE]int
	l int
	r int
}

func NewQueue() *Queue {
	q := &Queue{
		l: 0,
		r: -1,
	}
	return q
}

func (q *Queue) Enqueue(x int) {
	if q.r+1 >= MAX_SIZE {
		fmt.Println("Queue is full")
		return
	}
	q.r++
	q.a[q.r] = x
}

func (q *Queue) Dequeue() {
	if q.l > q.r {
		fmt.Println("Queue is empty")
		return
	}
	q.l++
}

func (q *Queue) Front() int {
	if q.l > q.r {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.a[q.l]
}

func (q *Queue) Size() int {
	return q.r - q.l + 1
}

func main() {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(6)
	q.Enqueue(7)
	q.Enqueue(9)
	q.Enqueue(16)
	q.Enqueue(16) // "Queue is full"

	fmt.Print(q.Front(), " ")
	fmt.Print(q.Front(), " ")
	fmt.Print(q.Front(), " ")
	fmt.Print(q.Front(), " ")
	fmt.Print(q.Front(), " ")
	q.Dequeue()
	fmt.Print(q.Front(), " ")
}
