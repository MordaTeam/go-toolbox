package main

import (
	"fmt"
	"log"

	"github.com/MordaTeam/go-toolbox/datastructs/priority_queue"
)

func main() {
	pq := priority_queue.New[string]()

	pq.Push(priority_queue.Node[string]{Payload: "foo", Priority: 5})
	pq.Push(priority_queue.Node[string]{Payload: "buz", Priority: 7})
	pq.Push(priority_queue.Node[string]{Payload: "bar", Priority: 10})
	pq.Push(priority_queue.Node[string]{Payload: "baz", Priority: 7})

	printQ(pq)

	for range 4 {
		el, err := pq.Pop()
		fatalOnerr(err)
		fmt.Println("popped", el)
		printQ(pq)
	}
}

func printQ[T any](pq priority_queue.PriorityQueue[T]) {
	for _, item := range pq.Items() {
		fmt.Printf("%2d: %+v\n", item.Priority, item.Payload)
	}
	fmt.Println("-----------------------")
}

func fatalOnerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
