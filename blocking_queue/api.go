package blocking_queue

type BlockingQueue interface {
	Push(item interface{})
	Pop() interface{}
}
