package main

import (
	"fmt"
	"github.com/NoobMaster-96/bounded-blocking-queue/blocking_queue"
	v2 "github.com/NoobMaster-96/bounded-blocking-queue/blocking_queue/v2"
	"time"
)

func main() {
	bQueue := v2.NewBoundedBlockingQueue(1)
	go pushToQueue1(bQueue)
	go pushToQueue2(bQueue)
	time.Sleep(2 * time.Second)
	go popQueue(bQueue)
	go popQueue(bQueue)
	time.Sleep(2 * time.Second)
}

func pushToQueue1(bQueue blocking_queue.BlockingQueue) {
	bQueue.Push(1)

}

func pushToQueue2(bQueue blocking_queue.BlockingQueue) {
	bQueue.Push(2)
}

func popQueue(bQueue blocking_queue.BlockingQueue) {
	fmt.Println("Item popped from the queue ", bQueue.Pop())
}
