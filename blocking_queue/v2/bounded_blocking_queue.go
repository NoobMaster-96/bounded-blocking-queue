package v2

import (
	"fmt"
	"sync"
)

type BoundedBlockingQueue struct {
	lock      sync.Mutex
	condition sync.Cond
	data      []interface{}
	capacity  int
}

func NewBoundedBlockingQueue(capacity int) *BoundedBlockingQueue {
	boundedBlockingQueue := new(BoundedBlockingQueue)
	boundedBlockingQueue.condition = sync.Cond{
		L: &boundedBlockingQueue.lock,
	}
	boundedBlockingQueue.capacity = capacity
	return boundedBlockingQueue
}

func (b *BoundedBlockingQueue) Push(item interface{}) {
	b.condition.L.Lock()
	defer b.condition.L.Unlock()

	for b.isFull() {
		b.condition.Wait()
	}

	b.data = append(b.data, item)
	fmt.Println("Item pushed to the queue ", item)
	b.condition.Signal()
}

func (b *BoundedBlockingQueue) Pop() interface{} {
	b.condition.L.Lock()
	defer b.condition.L.Unlock()

	for b.isEmpty() {
		b.condition.Wait()
	}

	result := b.data[0]
	b.data = b.data[1:len(b.data)]
	b.condition.Signal()
	return result
}

func (b *BoundedBlockingQueue) isFull() bool {
	return len(b.data) == b.capacity
}

func (b *BoundedBlockingQueue) isEmpty() bool {
	return len(b.data) == 0
}
