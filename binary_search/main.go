package main

import (
	"fmt"

	"example.com/demo/queue"
	"example.com/demo/stack"
)

// interview question
// time complexity O(log n) always
// space complexity O(1)
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 8, 9}
	r := binarySearch(arr, 6)
	fmt.Println("binary search", r)

	fmt.Println("-------------------------------------------------------")

	//queue
	q := queue.Queue{}
	fmt.Println("queue: ", q, "empty: ", q.IsEmpty())

	q.Enqueue(1)
	q.Enqueue(1)
	q.Enqueue(12)
	q.Enqueue(12)
	q.Enqueue(13)
	q.Enqueue(13)

	fmt.Println("queue: ", q, "empty: ", q.IsEmpty())

	v, err := q.Peek()
	fmt.Println("queue peek: ", q, "value: ", v, "err: ", err)

	v, err = q.Dequeue()
	fmt.Println("queue dequeue: ", q, "value: ", v, "err: ", err)

	fmt.Println("-------------------------------------------------------")

	//stack
	s := stack.Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(4)
	s.Push(4)
	s.Push(5)

	fmt.Println("Stack: ", s)
	v, err = s.Pop()
	fmt.Println("Stack POP: ", v, err)
	fmt.Println("Stack: ", s)

}
