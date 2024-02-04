package v1

import "fmt"

type BoundedBlockingQueue struct {
	channel chan interface{}
}

func NewBoundedBlockingQueue(capacity int) *BoundedBlockingQueue {
	return &BoundedBlockingQueue{
		channel: make(chan interface{}, capacity),
	}
}

func (b *BoundedBlockingQueue) Push(item interface{}) {
	b.channel <- item
	fmt.Println("Item pushed to the queue ", item)
	return
}

func (b *BoundedBlockingQueue) Pop() interface{} {
	return <-b.channel
}

