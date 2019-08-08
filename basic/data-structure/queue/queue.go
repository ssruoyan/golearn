package queue

import "sync"

type Queue struct {
	items []interface{}
	lock sync.RWMutex
}

func (queue *Queue) Push(data interface{}) {
	queue.lock.Lock()
	queue.items = append(queue.items, data)
	queue.lock.Unlock()
}

func (queue *Queue) Pop() interface{} {
	queue.lock.Lock()
	item := queue.items[0]
	queue.items = queue.items[1:len(queue.items)]
	queue.lock.Unlock()

	return &item
}

func (queue Queue) Size() int {
	return len(queue.items)
}