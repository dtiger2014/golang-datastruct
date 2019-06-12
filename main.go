package main

import (
	"fmt"
	"golang-datastruct/util/array"
	"golang-datastruct/util/arrayqueue"
	"golang-datastruct/util/arraystack"
	"golang-datastruct/util/loopqueue"
)

func main() {
	// TestArray()
	// TestMyArrayStack()
	// TestArrayQueue()
	TestLoopQueue()

	// var v interface{}
	// v = 1
	// switch t := v.(type) {
	// case int:
	// 	fmt.Println("int")
	// case float64:
	// 	fmt.Println("float64")
	// //... etc
	// default:
	// 	_ = t
	// 	fmt.Println("unkown")
	// }
}

func TestLoopQueue() {
	q := loopqueue.New()
	fmt.Println(q)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	fmt.Println(q)

	fmt.Println(q.GetFront())
}

func TestArrayQueue() {
	q := arrayqueue.New()

	fmt.Println(q)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	fmt.Println(q)

	fmt.Println(q.GetFront())

}

func TestArrayStack() {
	stack := arraystack.New()

	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	e, err := stack.Peek()
	fmt.Println(e)
	fmt.Println(err)
	fmt.Println(stack)
}

func TestArray() {
	arr := array.New()

	for i := 0; i < 11; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	arr.Add(1, 100)
	arr.AddFirst(200)
	arr.AddLast(300)
	fmt.Println(arr)

	fmt.Println(arr.Get(3))
	fmt.Println(arr.GetFirst())
	fmt.Println(arr.GetLast())

	arr.Set(1, 500)
	fmt.Println(arr)
	fmt.Println(arr.Contains(310))
	fmt.Println(arr.Find(3010))

	arr.Remove(1)
	arr.RemoveFirst()
	arr.RemoveLast()
	fmt.Println(arr)

	arr.RemoveElement(100)
	fmt.Println(arr)

	for i := 0; i < 6; i++ {
		arr.RemoveLast()
	}
	fmt.Println(arr)

	for i := 0; i < 20; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)
}
